package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	Requests             prometheus.CounterVec
	FunctionCount        prometheus.CounterVec
	FunctionFailure      prometheus.CounterVec
	FunctionResponseTime prometheus.GaugeVec
}

func (m *Metrics) AddRequest(endpoint, method string) {
	m.Requests.With(prometheus.Labels{"endpoint": endpoint, "method": method}).Add(1)
}

func (m *Metrics) AddFunctionCount(functionName string, failed bool) {
	if failed {
		m.FunctionFailure.With(prometheus.Labels{"function": functionName}).Add(1)
	} else {
		m.FunctionCount.With(prometheus.Labels{"function": functionName}).Add(1)
	}
}

func (m *Metrics) AddFunctionResponseTime(functionName string, since time.Time) {
	m.FunctionResponseTime.With(prometheus.Labels{"function": functionName}).Set(float64(time.Since(since) / 1000000))
}

// Register metrics, creates prometheus metrics for ghoster.
func Register(namespace, subsystem string) Metrics {
	return Metrics{
		Requests: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_requests",
			Help:      "getting total number of requests per endpoint",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"endpoint", "method"}),
		FunctionCount: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_function_calls",
			Help:      "getting total number of function calls",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
		FunctionFailure: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_function_failuers",
			Help:      "getting total number of function failure calls",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
		FunctionResponseTime: *promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "function_response_time",
			Help:      "response time of functions",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
	}
}
