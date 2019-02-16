package metrics

import "github.com/prometheus/client_golang/prometheus"

var Latency *prometheus.HistogramVec

func init() {
	Latency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "http_request_duration_seconds",
		Help:        "How long it took to process the request, partitioned by status code, method, HTTP path, platform and Browser name",
		ConstLabels: prometheus.Labels{"service": "IBM-iX-hack-talks-sample"},
	},
		[]string{"code", "method", "path", "platform", "browser"},
	)
	prometheus.MustRegister(Latency)
}
