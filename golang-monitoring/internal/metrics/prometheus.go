package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// count total HTTP requests
	TotalHTTPRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"code", "method", "endpoint"})

	// track HTTP request duration
	HTTPRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_requests_duration_seconnds",
		Help:    "HTTP request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "endpoint"})

	// track active connections
	ActiveConnections = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_connections",
		Help: "Number of active connections",
	})

	// track database operations
	DBOps = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "database_operations_total",
		Help: "Total number of database operations",
	}, []string{"operation", "status"})

	// track Redis operations
	RedisOps = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "redis_operations_total",
		Help: "Total number of Redis operations",
	}, []string{"operation", "status"})
)

// register all metrics
func Init() {
	prometheus.MustRegister(TotalHTTPRequests)
	prometheus.MustRegister(HTTPRequestDuration)
	prometheus.MustRegister(ActiveConnections)
	prometheus.MustRegister(DBOps)
	prometheus.MustRegister(RedisOps)
}
