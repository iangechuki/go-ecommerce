package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/iangechuki/go-ecommerce/internal/mailer"
	"github.com/iangechuki/go-ecommerce/internal/store"
)

type RegisterUserPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Username string `json:"username" validate:"required,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}
type UserTokenResponse struct {
	*store.User
	Token string
}

// registerUserHandler godoc
// @Summary      Register a user
// @Description  Register a user
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        payload body RegisterUserPayload true "RegisterUserPayload"
// @Success      200
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /auth/register [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
	}
	ctx := r.Context()

	user := &store.User{
		Email:      payload.Email,
		Username:   payload.Username,
		IsVerified: false,
	}
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.exp); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// send email with invite token(plain token)
	verificationLink := fmt.Sprintf("%s/verify?token=%s", app.config.frontendURL, plainToken)
	isProdEnv := app.config.env == "production"
	vars := struct {
		Username       string
		ActivationLink string
	}{
		Username:       user.Email,
		ActivationLink: verificationLink,
	}
	// instead of sending an email synchronously, we can send it asynchronously by queuing it as a
	// background job in a separate goroutine
	emailJob := Job{
		Type: "sendEmail",
		Payload: EmailPayload{
			Template: mailer.UserWelcomeTemplate,
			Username: user.Username,
			Email:    user.Email,
			Vars:     vars,
			Debug:    !isProdEnv,
			UserID:   user.ID,
			Ctx:      ctx,
		},
	}
	app.jobQueue <- emailJob
	if err := app.jsonResponse(w, http.StatusCreated, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

// verifyUserHandler godoc
// @Summary      Verify a user
// @Description  Verify a user
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        token query string true "Verification token"
// @Success      204
// @Failure      400 {object} error
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /auth/verify [get]
func (app *application) verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	log.Println(tokenString)
	if tokenString == "" {
		app.badRequestResponse(w, r, fmt.Errorf("token is required"))
		return
	}

	ctx := r.Context()
	err := app.store.Users.Verify(ctx, tokenString)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.notFoundResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}
type LoginResponse struct {
	*store.User
	AccessToken string
}

// loginUserHandler godoc
// @Summary      Login a user
// @Description  Login a user
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        payload body LoginPayload true "LoginPayload"
// @Success      200 {object} LoginResponse
// @Failure      400 {object} error
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /auth/login [post]
func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
	}
	ctx := r.Context()
	user, err := app.store.Users.GetByEmail(ctx, payload.Email)

	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid credentials"), "invalid credentials")
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if err := user.Password.Compare(payload.Password); err != nil {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid credentials"), "invalid credentials")
		return
	}
	if user.TwoFAEnabled {
		tempToken, err := app.Generate2FAToken(user.ID)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}
		response := struct {
			Requires2FA bool   `json:"requires_2fa"`
			Token       string `json:"token"`
		}{
			Requires2FA: true,
			Token:       tempToken,
		}
		app.jsonResponse(w, http.StatusOK, response)
		return

	}
	clientFingerprint := app.CreateClientFingerprint(r)
	// check for existing session from the same device
	existingSession, err := app.store.Sessions.GetByUserFingerprint(ctx, user.ID, clientFingerprint)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			break
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if existingSession != nil {
		if time.Since(existingSession.LastAccessed) < 5*time.Minute {
			app.badRequestResponse(w, r, fmt.Errorf("please wait before requesting a new token"))
			return
		}
		if err := app.store.Sessions.Delete(ctx, existingSession.ID); err != nil {
			app.internalServerError(w, r, err)
			return
		}
	}
	accessToken, err := app.GenerateAccessToken(user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	refreshToken, err := app.GenerateRefreshToken(user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	tokenHash := sha256.Sum256([]byte(refreshToken))
	hashString := hex.EncodeToString(tokenHash[:])
	session := &store.Session{
		UserID:            user.ID,
		TokenHash:         hashString,
		CreatedAt:         time.Now(),
		ExpiresAt:         time.Now().Add(app.config.auth.refreshToken.exp),
		IPAddress:         r.RemoteAddr,
		UserAgent:         r.UserAgent(),
		ClientFingerprint: clientFingerprint,
		LastAccessed:      time.Now(),
	}
	if err := app.store.Sessions.Create(ctx, session); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(app.config.auth.refreshToken.exp),
		MaxAge:   int(app.config.auth.refreshToken.exp.Seconds()),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   app.config.env == "production",
	}
	http.SetCookie(w, cookie)
	response := LoginResponse{User: user, AccessToken: accessToken}

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
	}
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// refreshTokenHandler godoc
// @Summary      Refresh a token
// @Description  Refresh a token
// @Tags authentication
// @Accept       json
// @Produce      json
// @Success      200 {object} RefreshTokenResponse
// @Failure      401 {object} error
// @Failure      500 {object} error
// Security     ApiKeyAuth
// @Router       /auth/refresh [get]
func (app *application) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")

	if err != nil {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("unauthorized"), "token is missing")
		return
	}
	refreshToken := cookie.Value

	ctx := r.Context()
	session, err := app.store.Sessions.GetByToken(ctx, refreshToken)

	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.unauthorizedErrorResponse(w, r, err, "")
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if session.ExpiresAt.Before(time.Now()) {
		app.unauthorizedErrorResponse(w, r, err, "session expired")
		// writeJSONError(w, http.StatusUnauthorized, "session expired")
		return
	}
	newAccessToken, err := app.GenerateAccessToken(session.UserID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.store.Sessions.UpdateLastAccessed(ctx, session.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	response := RefreshTokenResponse{AccessToken: newAccessToken}

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
	}
}

// logoutUserHandler godoc
// @Summary      Logout a user
// @Description  Logout a user
// @Tags authentication
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      401 {object} error
// @Failure      500 {object} error
// @Router       /auth/logout [get]
func (app *application) logoutUserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		app.unauthorizedErrorResponse(w, r, err, "token is missing")
		return
	}
	refreshToken := cookie.Value
	tokenHash := sha256.Sum256([]byte(refreshToken))
	hashString := hex.EncodeToString(tokenHash[:])
	ctx := r.Context()
	session, err := app.store.Sessions.GetByToken(ctx, hashString)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.unauthorizedErrorResponse(w, r, err, "")
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if err := app.store.Sessions.Delete(ctx, session.ID); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   app.config.env == "production",
	})

	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}
