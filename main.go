package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/agrea/watchtower/config"
	"github.com/agrea/watchtower/httpapi"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"port": os.Getenv("PORT"),
	}).Info("Starting Watchtower")

	r := chi.NewRouter()

	// HTTP middleware configuration.
	r.Use(NewStructuredLogger(logger))

	// HTTP endpoints.
	r.Mount("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(config.StaticPath))))
	r.Get("/", http.HandlerFunc(httpapi.Index))

	// Start the web server.
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
	if err != nil {
		logger.WithError(err).Fatal("A web server error occurred")
	}
}
