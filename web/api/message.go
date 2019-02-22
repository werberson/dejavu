package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	uagent "github.com/mileusna/useragent"
	"github.com/werberson/dejavu/web/metrics"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	// start latency time
	start := time.Now()

	ua := uagent.Parse(r.Header.Get("User-Agent"))
	browser := ua.Name
	platform := ua.OS

	if _, ok := buggedPlatform[strings.ToLower(platform)]; ok {
		time.Sleep(5 * time.Second)
	}

	statusCode := http.StatusOK
	if _, err := io.WriteString(w, fmt.Sprintf(`{"message": "Acessando do %v no sistema %v."}`, browser, platform)); err != nil {
		statusCode = http.StatusInternalServerError
		log.Print("Error occurred when writing response message", err)
	}

	w.Header().Set("Content-Type", "application/json")

	// register latency observation
	metrics.Latency.WithLabelValues(strconv.Itoa(statusCode), r.Method, r.URL.Path, platform, browser).
		Observe(time.Since(start).Seconds())
}
