package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
)

func init() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(httpRequestsTotal)
}

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
	}()
	w.Write([]byte("Hello, world!"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		httpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, "500").Inc()
	}()
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe(":3333", nil)
}
