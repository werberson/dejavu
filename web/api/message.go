package api

import (
	"fmt"
	"github.com/mssola/user_agent"
	"github.com/werberson/prometheus-metrics-sample/web/metrics"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	// start latency time
	start := time.Now()

	ua := user_agent.New(r.Header.Get("User-Agent"))
	browser, _ := ua.Browser()
	platform := ua.Platform()

	if _, ok := buggedPlatform[strings.ToLower(platform)]; ok {
		time.Sleep(5 * time.Second)
	}

	statusCode := http.StatusOK
	if _, err := io.WriteString(w, fmt.Sprintf(`{"message": "Hello %v on %v device."}`, browser, platform)); err != nil {
		statusCode = http.StatusInternalServerError
		log.Print("Error occurred when writing response message", err)
	}

	w.Header().Set("Content-Type", "application/json")

	// register latency observation
	metrics.Latency.WithLabelValues(strconv.Itoa(statusCode), r.Method, r.URL.Path, platform, browser).
		Observe(time.Since(start).Seconds())
}
