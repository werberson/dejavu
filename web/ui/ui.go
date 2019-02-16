package ui

import (
	"net/http"
)

func Handler() http.Handler {
	return http.FileServer(http.Dir("/go/src/github.com/werberson/prometheus-metrics-sample/web/ui/static/"))
}
