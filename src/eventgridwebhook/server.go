package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Microsoft/go-eventgrid-handler/eventgrid"
	"github.com/Microsoft/go-eventgrid-handler/logb"
	"github.com/gorilla/mux"
)

// channel used to send os.Interrupts
var osChan = make(chan os.Signal, 1)

// runServer - setup and run the web server
// this blocks until the web server exits
func runServer(port int) error {

	// use gorilla mux
	r := mux.NewRouter()

	// this is our only handler
	// chain the handlers together as middleware
	// app services does https offloading, so check for the x-forwarded-proto header
	// only accept POST requests
	// change the URI based on your webhook definition
	r.Handle("/person", logb.Handler(eventgrid.Handler(personHandler))).Methods("POST").Headers("x-forwarded-proto", "https")
	http.Handle("/", r)

	// setup the server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(port),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// run webserver in a Go routine so we can cancel
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("ERROR:", err)
		}
	}()

	// Block until we receive our signal
	signal.Notify(osChan, os.Interrupt)
	<-osChan

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	return srv.Shutdown(ctx)
}
