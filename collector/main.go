package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const webPort = "80"

func main() {
	log.Printf("Starting Collector on port %s\n", webPort)

	_, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: NewCollector().Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
