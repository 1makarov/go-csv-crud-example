package main

import (
	"context"
	"github.com/1makarov/go-csv-crud-example/internal/config"
	"github.com/1makarov/go-csv-crud-example/internal/db/postgres"
	"github.com/1makarov/go-csv-crud-example/internal/pkg/signaler"
	"github.com/1makarov/go-csv-crud-example/internal/repository"
	"github.com/1makarov/go-csv-crud-example/internal/server"
	"github.com/1makarov/go-csv-crud-example/internal/services"
	"github.com/1makarov/go-csv-crud-example/internal/transport"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sirupsen/logrus"
)

// @title Store API
// @version 1.0

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg := config.Init()

	db, err := postgres.Open(cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
	}

	repo := repository.New(db)
	service := services.New(repo)
	handler := transport.NewHandler(service)

	srv := server.New(cfg.HTTP.Port, handler.Init(cfg))

	go func() {
		if err = srv.Run(); err != nil {
			logrus.Errorf("error start server: %s", err.Error())
		}
	}()

	logrus.Println("rest-api started")

	signaler.Wait()

	if err = srv.Stop(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
