package main

import (
	"flag"
	"log"
	"os"
)

type config struct {
	port int
	env  string
}

type application struct {
	config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "api server port")
	flag.StringVar(&cfg.env, "environment", "development", "Application environment")

	flag.Parse()

	logger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Fatal(err)
}
