package main

import (
	"log"
	"time"

	"github.com/iangechuki/go-ecommerce/internal/auth"
	"github.com/iangechuki/go-ecommerce/internal/db"
	"github.com/iangechuki/go-ecommerce/internal/env"
	"github.com/iangechuki/go-ecommerce/internal/mailer"
	"github.com/iangechuki/go-ecommerce/internal/store"
	"go.uber.org/zap"
)

//	@title	Go Ecommerce API

//	@description	API for Go ECommerce
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	config := config{
		addr: ":8080",
		env:  env.GetString("ENV", "development"),

		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/go-ecommerce?sslmode=disable"),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 10),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		apiURL:      env.GetString("API_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
		auth: authConfig{
			accessToken: tokenConfig{
				secret: env.GetString("ACCESS_TOKEN_SECRET", "secret"),
				exp:    time.Minute * 15,
			},
			refreshToken: tokenConfig{
				secret: env.GetString("REFRESH_TOKEN_SECRET", "secret"),
				exp:    time.Hour * 24 * 7,
			},
			iss: env.GetString("ISS", "go-ecommerce"),
			aud: env.GetString("AUD", "go-ecommerce"),
		},
		mail: mailConfig{
			fromEmail: env.GetString("FROM_EMAIL", ""),
			resend: resendConfig{
				apiKey: env.GetString("RESEND_API_KEY", ""),
			},
			exp: time.Hour * 24 * 3,
		},
	}
	// Logger
	logger := zap.Must(zap.NewDevelopment()).Sugar()
	defer logger.Sync()

	// database
	db, err := db.New(
		config.db.addr,
		config.db.maxOpenConns,
		config.db.maxIdleConns,
		config.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := store.NewPostgresStorage(db)
	accessAuthenticator := auth.NewJWTAuthenticator(
		config.auth.accessToken.secret,
		config.auth.aud,
		config.auth.iss,
	)
	refreshAuthenticator := auth.NewJWTAuthenticator(
		config.auth.refreshToken.secret,
		config.auth.aud,
		config.auth.iss,
	)

	// Mailer
	resendClient, err := mailer.NewResendClient(config.mail.fromEmail, config.mail.resend.apiKey)
	if err != nil {
		log.Fatal(err)
	}
	app := application{
		config:               config,
		store:                store,
		accessAuthenticator:  accessAuthenticator,
		refreshAuthenticator: refreshAuthenticator,
		mailer:               resendClient,
		logger:               logger,
		jobQueue:             make(chan Job, 100),
	}
	app.startJobWorker()
	app.ScheduleExpiredSessionCleanup()
	log.Println("name", env.GetString("FROM_EMAIL", "test"))
	mux := app.mount()
	log.Fatal(app.run(mux))
}
