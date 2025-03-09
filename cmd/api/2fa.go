package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iangechuki/go-ecommerce/internal/store"
	"github.com/pquerna/otp/totp"
)

type Enable2FAResponse struct {
	Secret string `json:"secret"`
	QRCode string `json:"qrcode"`
}

// enable2FAHandler godoc
// @Summary      Enable 2FA
// @Description  Enable 2FA
// @Tags 2FA
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      401 {object} error
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /user/2fa/enable [get]
func (app *application) enable2FAHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := getUserFromContext(ctx)
	fmt.Println("user", user)
	if user == nil {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("unauthorized"), "authentication is required")
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      app.config.auth.iss,
		AccountName: user.Email,
		SecretSize:  20,
	})
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	response := Enable2FAResponse{
		Secret: key.Secret(),
		QRCode: key.URL(),
	}
	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
func (app *application) disable2FAHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type Verify2FAPayload struct {
	Secret string `json:"secret" validate:"required"`
	Code   string `json:"code" validate:"required"`
}

// verify2FAHandler godoc
// @Summary      Verify 2FA
// @Description  Verify 2FA
// @Tags 2FA
// @Accept       json
// @Produce      json
// @Param        payload body Verify2FAPayload true "Verify2FAPayload"
// @Success      200
// @Failure      401 {object} error
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /user/2fa/verify [post]
func (app *application) verify2FAHandler(w http.ResponseWriter, r *http.Request) {
	var payload Verify2FAPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	user := getUserFromContext(r.Context())
	if user == nil {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("unauthorized"), "authentication is required")
		return
	}
	valid := totp.Validate(payload.Code, payload.Secret)
	if !valid {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid 2fa code"), "invalid 2fa code")
		return
	}
	user.TwoFASecret = payload.Secret
	user.TwoFAEnabled = true
	if err := app.store.Users.Update(r.Context(), user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusOK, "two enabled sucessfully"); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type Login2FAPayload struct {
	Code  string `json:"code" validate:"required,min=6,max=6"`
	Token string `json:"token" validate:"required"`
}

// login2FAHandler godoc
// @Summary      Login with 2FA
// @Description  Verify 2FA code and issue tokens
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        payload body Login2FAPayload true "Login2FAPayload"
// @Success      200 {object} LoginResponse
// @Failure      400 {object} error
// @Failure      401 {object} error
// @Failure      500 {object} error
// @Router       /auth/2fa/login [post]
func (app *application) login2FAHandler(w http.ResponseWriter, r *http.Request) {
	var payload Login2FAPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	jwtToken, err := app.accessAuthenticator.ValidateToken(payload.Token)
	if err != nil {
		app.unauthorizedErrorResponse(w, r, err, "invalid token")
		return
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		app.unauthorizedErrorResponse(w, r, err, "unauthorized")
		return
	}
	ctx := r.Context()
	user, err := app.getUser(ctx, userID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	valid := totp.Validate(payload.Code, user.TwoFASecret)
	if !valid {
		app.unauthorizedErrorResponse(w, r, fmt.Errorf("invalid 2fa code"), "invalid 2fa code")
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
