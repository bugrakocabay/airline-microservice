package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bugrakocabay/airline/auth-service/domain/api"
	"github.com/bugrakocabay/airline/auth-service/domain/config"
	"github.com/bugrakocabay/airline/auth-service/domain/token"
	"github.com/bugrakocabay/airline/auth-service/infrastructure/pg_repository"

	"github.com/jackc/pgx/v4"
)

const webPort = "80"

func main() {
	log.Printf("Starting Auth Service on port %s\n", webPort)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	serviceName := "dice"
	serviceVersion := "0.1.0"
	otelShutdown, err := config.SetupOTelSDK(ctx, serviceName, serviceVersion)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Setup Postgres
	pgConn, err := pgx.Connect(context.Background(), os.Getenv("DSN"))
	if err != nil {
		log.Panicf("error connecting postgres: %s", err)
	}

	// Setup token maker
	tokenMaker, err := token.NewPasetoMaker(os.Getenv("SYMMETRIC_KEY"))
	pgUserRepo := &pg_repository.PgUserRepository{Conn: pgConn}
	authHandler := api.NewAuthHandler(pgUserRepo, tokenMaker)

	// Setup HTTP Server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: authHandler.Routes(),
	}

	// Start HTTP Server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
