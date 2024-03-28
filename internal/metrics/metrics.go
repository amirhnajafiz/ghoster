package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	ListRequests    prometheus.Counter
	ExecuteRequests prometheus.CounterVec
	FunctionFailure prometheus.CounterVec
}

func Register() Metrics {
	return Metrics{
		ListRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: "total_list_requests",
			Help: "getting total number of list requests",
		}),
		ExecuteRequests: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "total_function_calls",
			Help: "getting total number of function calls",
		}, []string{"function"}),
		FunctionFailure: *promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "total_function_failuers",
			Help: "getting total number of function failure calls",
		}, []string{"function"}),
	}
}
