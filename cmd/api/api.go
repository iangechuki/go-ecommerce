package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iangechuki/go-ecommerce/docs"
	"github.com/iangechuki/go-ecommerce/internal/auth"
	"github.com/iangechuki/go-ecommerce/internal/mailer"
	"github.com/iangechuki/go-ecommerce/internal/store"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"go.uber.org/zap"
)

type application struct {
	config               config
	store                store.Storage
	logger               *zap.SugaredLogger
	accessAuthenticator  auth.Authenticator
	refreshAuthenticator auth.Authenticator
	mailer               mailer.Client
}

type config struct {
	addr        string
	env         string
	apiURL      string
	frontendURL string
	db          dbConfig
	auth        authConfig
	mail        mailConfig
}
type dbConfig struct {
	addr         string
	maxIdleConns int
	maxOpenConns int
	maxIdleTime  string
}
type authConfig struct {
	accessToken  tokenConfig
	refreshToken tokenConfig
	iss          string
	aud          string
}

type tokenConfig struct {
	secret string
	exp    time.Duration
}
type mailConfig struct {
	fromEmail string
	exp       time.Duration
	resend    resendConfig
}
type resendConfig struct {
	apiKey string
}

// HealthCheck godoc
// @Summary      Health check
// Description  HealthCheck url
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /health [get]
// @Failure      500 {object} error
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (app *application) mount() *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/v1/swagger/doc.json"), //The url pointing to API definition
		))
		r.Get("/health", healthCheckHandler)

		// r.Get("/mail/verify-account", app.previewVerifyAccountEmailHandler)
		// r.Get("/mail/verify-account/send", app.sendVerifyEmailTrial)
		// r.Get("/mail/reset-password", app.previewEmailForPasswordResetHandler)
		// r.Get("/mail/reset-password/send", app.sendPasswordResetEmailTrial)

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", app.registerUserHandler)
			r.Post("/login", app.loginUserHandler)
			r.Get("/verify", app.verifyUserHandler)
			r.Post("/forgot-password", app.forgotPasswordHandler)
			r.Post("/reset-password", app.resetPassword)

			r.Get("/refresh", app.refreshTokenHandler)
			r.Post("/logout", app.logoutUserHandler)
		})
		r.Route("/user", func(r chi.Router) {
			r.Post("/change-password", app.changePasswordHandler)
		})
		r.Route("/products", func(r chi.Router) {
			r.Get("/", app.listProductsHandler)
			// r.Post("/", app.createProductHandler)
		})
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Version = "0.0.0"
	docs.SwaggerInfo.BasePath = "/v1"
	srv := http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}
	log.Printf("server has started")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
