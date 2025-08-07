package main


import (
	"fmt"
	"log"
	"net/http"
	"time"
)


func serve (app *application) error  {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", app.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	log.Printf("Starting server on port %d", app.port)
	return server.ListenAndServe()
}