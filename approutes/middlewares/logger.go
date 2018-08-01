package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Log will create log files to the console
// accept string format and infinite number of args to create print format
func Log(format string, args ...interface{}) {
	// declare the outfile for the log
	// use os method from GoLang dugger method
	f, err := os.OpenFile("tmp/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)                 // create context
	log.Printf(format, args...)      // log to file
	fmt.Printf("> "+format, args...) // log to console
}

// Logger middleware to handle writing to external files and to the console
// accepts and returns http.Handler
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()                                               // start timer
		h.ServeHTTP(w, r)                                              // serve request
		t1 := time.Now()                                               // clock time to serve
		elapsed := t1.Sub(t0)                                          // initialize benchmark
		Log("[%v] - %s%s@%s \n", elapsed, r.Host, r.URL, r.RemoteAddr) // Log
	})
}
