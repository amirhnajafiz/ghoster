package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	requests             prometheus.CounterVec
	functionCount        prometheus.CounterVec
	functionFailure      prometheus.CounterVec
	functionResponseTime prometheus.GaugeVec
	pendingWorkers       prometheus.Gauge
}

func (m *Metrics) AddRequest(endpoint, method string) {
	m.requests.With(prometheus.Labels{"endpoint": endpoint, "method": method}).Add(1)
}

func (m *Metrics) AddWorker() {
	m.pendingWorkers.Add(1)
}

func (m *Metrics) RemoveWorker() {
	m.pendingWorkers.Sub(1)
}

func (m *Metrics) AddFunctionCount(functionName string, failed bool) {
	if failed {
		m.functionFailure.With(prometheus.Labels{"function": functionName}).Add(1)
	} else {
		m.functionCount.With(prometheus.Labels{"function": functionName}).Add(1)
	}
}

func (m *Metrics) AddFunctionResponseTime(functionName string, since time.Duration) {
	m.functionResponseTime.With(prometheus.Labels{"function": functionName}).Set(float64((since / 1000000)))
}

// Register metrics, creates prometheus metrics for ghoster
func Register(namespace, subsystem string) Metrics {
	return Metrics{
		requests: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_requests",
			Help:      "getting total number of requests per endpoint",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"endpoint", "method"}),
		functionCount: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_function_calls",
			Help:      "getting total number of function calls",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
		functionFailure: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name:      "total_function_failuers",
			Help:      "getting total number of function failure calls",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
		functionResponseTime: *promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name:      "function_response_time",
			Help:      "response time of functions",
			Namespace: namespace,
			Subsystem: subsystem,
		}, []string{"function"}),
		pendingWorkers: promauto.NewGauge(prometheus.GaugeOpts{
			Name:      "pending_workers",
			Help:      "current number of pend workers",
			Namespace: namespace,
			Subsystem: subsystem,
		}),
	}
}
