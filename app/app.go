package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sshindanai/bookstore-utils-go/logger"
)

var (
	r = mux.NewRouter()
)

// StartApp is the entry point of the app.
func StartApp() {
	mapUrls()

	srv := &http.Server{
		Addr: ":8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	logger.Logger().Info("app is starting...")

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
