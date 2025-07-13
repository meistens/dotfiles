package main

import (
	"context"
	"log"
	"net/http"
	"skello/internal/cache"
	"skello/internal/config"
	"skello/internal/db"
	"skello/internal/handlers"
	"skello/internal/logger"
	"skello/internal/metrics"
	"skello/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger.Init()
	metrics.Init()

	db.MustInit(context.Background())
	defer db.Close()

	cache.Init()
	defer cache.Close()

	// create router
	ruter := mux.NewRouter()
	// add logging middleware
	ruter.Use(middleware.LoggingMiddleware)
	// health handler
	healthHandler := handlers.NewHealthHandler()
	ruter.HandleFunc("/health", healthHandler.Health).Methods("GET")
	// user handlers
	userHandler := handlers.NewUserHandler()
	ruter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	ruter.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	// metrics endpoint with promhttp
	ruter.Handle("/metrics", promhttp.Handler())

	port := config.GetEnv("PORT", "8080")
	logger.Get().Infof("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, ruter))
}
