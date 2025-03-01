package main

import (
	"log"

	"github.com/iangechuki/go-ecommerce/internal/db"
	"github.com/iangechuki/go-ecommerce/internal/env"
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

func main() {
	config := config{
		addr: ":8080",
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/go-ecommerce?sslmode=disable"),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 10),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		apiURL: env.GetString("API_URL", "localhost:8080"),
	}

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
	app := application{
		config: config,
	}
	log.Println("name", env.GetString("NAME", "test"))
	mux := app.mount()
	log.Fatal(app.run(mux))
}
