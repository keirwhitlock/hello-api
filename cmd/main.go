package main

import (
	"fmt"
	"github.com/keirwhitlock/hello-api/handlers"
	"github.com/keirwhitlock/hello-api/handlers/rest"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatalf("ERROR: %s", err)
}
