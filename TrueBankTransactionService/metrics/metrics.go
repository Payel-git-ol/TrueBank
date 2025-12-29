package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	KafkaMessagesOut = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kafka_messages_out_total",
		Help: "Total number of Kafka messages produced",
	})
	KafkaMessagesIn = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "kafka_messages_in_total",
		Help: "Total number of Kafka messages consumed",
	})
)

var (
	GRPCRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "status"},
	)

	GRPCRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
)

func Init() {
	prometheus.MustRegister(KafkaMessagesOut, KafkaMessagesIn)
}
