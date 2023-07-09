package main

import (
	"context"
	"fmt"
	"github.com/bugrakocabay/airline/auth-service/domain/token"
	"log"
	"net/http"
	"os"

	"github.com/bugrakocabay/airline/auth-service/domain/api"
	"github.com/bugrakocabay/airline/auth-service/infrastructure/pg_repository"

	"github.com/jackc/pgx/v4"
)

const webPort = "80"

func main() {
	log.Printf("Starting Auth Service on ports %s\n", webPort)

	pgConn, err := pgx.Connect(context.Background(), os.Getenv("DSN"))
	if err != nil {
		log.Panicf("error connecting postgres: %s", err)
	}

	tokenMaker, err := token.NewPasetoMaker(os.Getenv("SYMMETRIC_KEY"))
	pgUserRepo := &pg_repository.PgUserRepository{Conn: pgConn}
	authHandler := api.NewAuthHandler(pgUserRepo, tokenMaker)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: authHandler.Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
