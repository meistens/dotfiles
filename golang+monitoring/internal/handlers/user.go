package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"skello/internal/cache"
	"skello/internal/db"
	"skello/internal/logger"
	"skello/internal/metrics"
	"skello/internal/models"
	"time"

	"github.com/sirupsen/logrus"
)

// UserHandler handles user-related requests
type UserHandler struct{}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// get context from request
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Get().WithError(err).Error("Failed to decode user JSON")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// insert user to db
	err := db.Get().QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)

	if err != nil {
		logger.Get().WithError(err).Error("Failed to insert user into the database")
		metrics.DBOps.WithLabelValues("insert", "error").Inc()
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	metrics.DBOps.WithLabelValues("insert", "success").Inc()

	// cache, this scenario is the users
	userJSON, _ := json.Marshal(user)
	cacheKey := fmt.Sprintf("user:%d", user.ID)
	err = cache.Get().Set(ctx, cacheKey, userJSON, time.Hour).Err()

	if err != nil {
		logger.Get().WithError(err).Error("Failed to cache user in Redis")
		metrics.RedisOps.WithLabelValues("set", "error").Inc()
	} else {
		metrics.RedisOps.WithLabelValues("set", "success")
	}

	logger.Get().WithFields(logrus.Fields{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
	}).Info("User created successfully")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
