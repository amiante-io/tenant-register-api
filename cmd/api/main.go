package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config config
	logger slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))

	app := &application{
		config: cfg,
		logger: *logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
