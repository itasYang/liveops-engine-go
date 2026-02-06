package main

import (
	"log"
	"net/http"
	"time"

	apphttp "github.com/itasYang/liveops-engine-go/internal/http"
)

func main() {
	srv := apphttp.NewServer()

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           srv.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("api listening on %s", httpServer.Addr)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
