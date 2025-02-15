package main

import (
	"flag"
	"fmt"
	"github.com/bhusal-rj/design-pattern/config"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *config.Application
}

type appConfig struct {
	useCache bool
}

func main() {
	godotenv.Load(".env")
	app := application{}

	//set default flag to false
	flag.BoolVar(&app.config.useCache, "cache", false, "Use template Cache")
	flag.Parse()

	//get database
	db, err := initMySQLDB()
	if err != nil {
		log.Panic(err)
	}
	app.App = config.New(db)
	//create the server instance
	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port", port)

	//listen and create the server instance
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
