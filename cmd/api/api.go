package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iangechuki/go-ecommerce/docs"
	_ "github.com/iangechuki/go-ecommerce/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type application struct {
	config config
}

type config struct {
	addr   string
	apiURL string
	db     dbConfig
}
type dbConfig struct {
	addr         string
	maxIdleConns int
	maxOpenConns int
	maxIdleTime  string
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
