package main

import (
	"electronicstore/internal/data"
	db "electronicstore/internal/db"
	"electronicstore/internal/jsonlog"
	"electronicstore/pkg/config"
	"os"
	"sync"
	// Import the pq driver so that it can register itself with the database/sql
	// package. Note that we alias this import to the blank identifier, to stop the Go
	// compiler complaining that the package isn't being used.
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type application struct {
	config *config.Config
	logger *jsonlog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg = config.GetConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	db, err := db.OpenDB()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: &cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
