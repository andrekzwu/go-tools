// Copyright 2020-06-19 aweu
package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

// define process time enum
const (
	ProcessTimeWorse = 3000
	ProcessTimeBad   = 1000
	ProcessTimeOk    = 500
)

// define process time status enum
const (
	OkStatus = iota + 1
	BadStatus
	WorseStatus
)

var (
	apiRequest *prometheus.CounterVec
	apiCount   *prometheus.CounterVec
	apiTime    *prometheus.SummaryVec
)

var (
	defaultPattern = "/metrics"
	defaultAddress = "0.0.0.0:8081"
)

// define system config
type Config struct {
	NameSpace string
	System    string
}

// register and listen
func Register(config Config) {
	nameSpace := strings.Replace(config.NameSpace, "-", "_", -1)
	system := strings.Replace(config.System, "-", "_", -1)
	apiRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: nameSpace,
			Subsystem: system,
			Name:      "api_request",
			Help:      "api request counter",
		},
		[]string{"handleMethod"},
	)
	apiCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: nameSpace,
			Subsystem: system,
			Name:      "api_count",
			Help:      "api handler counter",
		},
		[]string{"handleMethod", "status"},
	)
	apiTime = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: nameSpace,
			Subsystem: system,
			Name:      "api_time",
			Help:      "api process time",
		},
		[]string{"handleMethod", "status"},
	)
	// Register
	prometheus.MustRegister(apiRequest)
	prometheus.MustRegister(apiCount)
	prometheus.MustRegister(apiTime)
	// Listen
	http.Handle(defaultPattern, promhttp.Handler())
	http.ListenAndServe(defaultAddress, nil)
}

// Record
// apiPath: api path, example：api/order/create/v1
// resultStatus: api response result status, example：0
// processTime:  api process time, example 200ms
func Record(apiPath string, resultStatus string, processTime int64) {
	apiCount.WithLabelValues(apiPath, resultStatus).Inc()
	apiTime.WithLabelValues(apiPath, fmt.Sprint(convertProcessTime2Status(processTime))).Observe(1)
}

// Count request
func CountRequest(apiPath string) {
	apiRequest.WithLabelValues(apiPath).Inc()
}

// processing time, in milliseconds
func convertProcessTime2Status(processTime int64) int {
	var status int
	if processTime > ProcessTimeWorse {
		status = WorseStatus
	} else if processTime > ProcessTimeBad {
		status = BadStatus
	} else if processTime > ProcessTimeOk {
		status = OkStatus
	}
	return status
}
