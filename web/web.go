package web

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/werberson/dejavu/web/api"
	"github.com/werberson/dejavu/web/ui"
)

func Initialize() error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/message", api.MessageHandler).Methods("GET")
	router.HandleFunc("/api/bug/{platform}", api.AddBugHandler).Methods("POST")
	router.HandleFunc("/api/bug/{platform}", api.RemoveBugHandler).Methods("DELETE")

	router.Handle("/", ui.Handler())
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Initialized")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("server initializing error", err)
		return err
	}
	return nil
}
