package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/iangechuki/go-ecommerce/internal/mailer"
	"github.com/iangechuki/go-ecommerce/internal/store"
)

type ForgotPassswordPayload struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

// forgotPasswordHandler godoc
//
// @Summary      Forgot password
// @Description  Forgot password
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        payload body ForgotPassswordPayload true "ForgotPassswordPayload"
// @Success      200
// @Failure      400 {object} error
// @Failure      500 {object} error
// @Router       /auth/forgot-password [post]
func (app *application) forgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var payload ForgotPassswordPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := r.Context()
	user, err := app.store.Users.GetByEmail(ctx, payload.Email)
	if err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.unauthorizedErrorResponse(w, r, err, "invalid credentials")
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	err = app.store.Users.CreatePasswordResetToken(ctx, user, hashToken, app.config.auth.accessToken.exp)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	passwordResetLink := fmt.Sprintf("%s/reset-password?token=%s", app.config.frontendURL, plainToken)
	vars := struct {
		Username          string
		PasswordResetLink string
	}{
		Username:          user.Username,
		PasswordResetLink: passwordResetLink,
	}
	sendID, err := app.mailer.Send(mailer.PasswordResetTemplate, user.Username, user.Email, vars, true)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	log.Printf("email sent with id %s", sendID)
	if err := app.jsonResponse(w, http.StatusOK, "password reset email sent"); err != nil {
		app.internalServerError(w, r, err)
	}
}

type ResetPasswordPayload struct {
	Password string `json:"password" validate:"required,min=6,max=255"`
}

// resetPassword godoc
//
// @Summary      Reset password
// @Description  Reset password
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        token query string true "Reset token"
// @Param        payload body ResetPasswordPayload true "ResetPasswordPayload"
// @Success      200
// @Failure      400 {object} error
// @Failure      500 {object} error
// @Router       /auth/reset-password [post]
func (app *application) resetPassword(w http.ResponseWriter, r *http.Request) {
	var payload ResetPasswordPayload
	token := r.URL.Query().Get("token")
	if token == "" {
		app.badRequestResponse(w, r, fmt.Errorf("token is required"))
		return
	}
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := r.Context()
	if err := app.store.Users.ResetPasswordUsingToken(ctx, token, payload.Password); err != nil {
		switch err {
		case store.ErrRecordNotFound:
			app.unauthorizedErrorResponse(w, r, err, "invalid token")
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}
	if err := app.store.Sessions.DeleteByUserID(ctx, 1); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusOK, "password reset successfully"); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old_password" validate:"required,min=6,max=255"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=255"`
	Email       string `json:"email" validate:"required,email,max=255"`
}

// changePasswordHandler godoc
//
// @Summary      Change password
// @Description  Change password
// @Tags authentication
// @Accept       json
// @Produce      json
// @Param        payload body ChangePasswordPayload true "ChangePasswordPayload"
// @Success      200
// @Failure      400 {object} error
// @Failure      500 {object} error
// @Security     ApiKeyAuth
// @Router       /user/change-password [post]
func (app *application) changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var payload ChangePasswordPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	ctx := r.Context()
	user, err := app.store.Users.GetByEmail(ctx, payload.Email)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := user.Password.Compare(payload.OldPassword); err != nil {
		app.unauthorizedErrorResponse(w, r, err, "invalid credentials")
		return
	}
	if err := user.Password.Set(payload.NewPassword); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.store.Users.Update(ctx, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
	if err := app.jsonResponse(w, http.StatusOK, "password changed successfully"); err != nil {
		app.internalServerError(w, r, err)
	}
}
