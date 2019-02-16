package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var buggedPlatform = make(map[string]string)

func AddBugHandler(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	platform := strings.ToLower(vars["platform"])
	buggedPlatform[platform] = platform
	log.Printf("Bug added for %s platform requests", platform)
}

func RemoveBugHandler(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	platform := strings.ToLower(vars["platform"])
	delete(buggedPlatform, platform)
	log.Printf("Bug removed for %s platform requests", platform)
}
