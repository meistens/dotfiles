package middleware

import (
	"fmt"
	"net/http"
	"skello/internal/logger"
	"skello/internal/metrics"
	"time"

	"github.com/sirupsen/logrus"
)

// respWrter wraps ResponseWriter to capture statusCode
type respWrter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *respWrter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs HTTP requests and tracks metrics
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		metrics.ActiveConnections.Inc()
		defer metrics.ActiveConnections.Dec()

		// create custom rspwrtr to capture status
		rw := &respWrter{ResponseWriter: w, statusCode: 200}

		// log request
		logger.Get().WithFields(logrus.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"remote_ip":  r.RemoteAddr,
			"user_agent": r.Header.Get("User-Agent"),
		}).Info("HTTP request started")

		next.ServeHTTP(rw, r)

		duration := time.Since(start).Seconds()
		statusCode := fmt.Sprintf("%d", rw.statusCode)

		// record metrics
		metrics.HTTPRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		metrics.TotalHTTPRequests.WithLabelValues(r.Method, r.URL.Path, statusCode).Inc()

		// Log response
		logger.Get().WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"status_code": rw.statusCode,
			"duration":    duration,
			"remote_ip":   r.RemoteAddr,
		}).Info("HTTP request completed")
	})
}
