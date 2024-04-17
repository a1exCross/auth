package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "auth_space"
	appName   = "auth_app"
)

var requestCounter prometheus.Counter
var responseCounter *prometheus.CounterVec
var histogramResponseTime *prometheus.HistogramVec

// Init - инициализация метрик
func Init(_ context.Context) error {
	requestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "grpc",
		Name:      appName + "_request_total",
		Help:      "Количество запросов к серверу gRPC",
	})

	responseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "grpc",
		Name:      appName + "_responses_total",
		Help:      "Количество ответов от сервера",
	}, []string{"status", "method"})

	histogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: "grpc",
		Name:      appName + "_histogram_response_time",
		Help:      "Время ответа оть сервера",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	}, []string{"status"})

	return nil
}

// IncRequestCounter - увеличивает счетчик метрики, показывающую количество запросов
func IncRequestCounter() {
	requestCounter.Inc()
}

// IncResponseCounter - увеличивает счетчик метрики, показывающую количество ответов
func IncResponseCounter(status, method string) {
	responseCounter.WithLabelValues(status, method).Inc()
}

// HistogramResponseTimeObserve - обновляет гистограмму врмени ответа от сервера
func HistogramResponseTimeObserve(status string, time float64) {
	histogramResponseTime.WithLabelValues(status).Observe(time)
}
