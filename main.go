package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/heptiolabs/healthcheck"
)

func main() {
	health := healthcheck.NewHandler()

	arguments := os.Args[1:]
	for _, element := range arguments {
		health.AddLivenessCheck(
			"service",
			healthcheck.TCPDialCheck(element, 5*time.Second),
		)
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:8086", health))
}
