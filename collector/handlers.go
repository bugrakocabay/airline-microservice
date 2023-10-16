package main

import (
	"log"
	"net/http"
)

type Collector struct {
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) handleCollection(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)

	w.WriteHeader(http.StatusOK)
}
