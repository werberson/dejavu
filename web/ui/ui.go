package ui

import (
	"net/http"
)

func Handler() http.Handler {
	return http.FileServer(http.Dir("./web/ui/static/"))
}
