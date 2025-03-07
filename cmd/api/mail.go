package main

import (
	"log"
	"net/http"

	"github.com/iangechuki/go-ecommerce/internal/mailer"
)

func (app *application) previewVerifyAccountEmailHandler(w http.ResponseWriter, r *http.Request) {
	sampleData := map[string]interface{}{
		"Username":      "Ian ochako",
		"ActivationURL": "http://localhost:3000/activate?token=1234",
	}
	subject, body, err := mailer.PreviewTemplate(mailer.UserWelcomeTemplate, sampleData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(subject))
	w.Write([]byte(body))
}
func (app *application) previewEmailForPasswordResetHandler(w http.ResponseWriter,
	r *http.Request) {
	sampleData := map[string]interface{}{
		"Username":          "Ian ochako",
		"PasswordResetLink": "http://localhost:3000/activate?token=1234",
	}
	subject, body, err := mailer.PreviewTemplate(mailer.PasswordResetTemplate, sampleData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(subject))
	w.Write([]byte(body))
}
func (app *application) sendVerifyEmailTrial(w http.ResponseWriter, r *http.Request) {
	sampleData := map[string]interface{}{
		"Username":      "Ian ochako",
		"ActivationURL": "http://localhost:3000/activate?token=1234",
	}
	// delivered@resend.dev
	sentId, err := app.mailer.Send(mailer.UserWelcomeTemplate, "Ian", "iangechuki@gmail.com", sampleData, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("email sent with id %s", sentId)
	jsonResponse(w, http.StatusOK, "email sent successfully")
}
func (app *application) sendPasswordResetEmailTrial(w http.ResponseWriter, r *http.Request) {
	sampleData := map[string]interface{}{
		"Username":          "Ian ochako",
		"PasswordResetLink": "http://localhost:3000/reset-password?token=1234",
	}
	// delivered@resend.dev
	sentId, err := app.mailer.Send(mailer.PasswordResetTemplate, "Ian", "iangechuki@gmail.com", sampleData, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("email sent with id %s", sentId)
	jsonResponse(w, http.StatusOK, "email sent successfully")
}
