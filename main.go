package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// HelloHandler return hello world message`
func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

// PingHandler Return message "Pong - " with current time in Month Day Time format
func PingHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	io.WriteString(w, fmt.Sprint("Pong - ", t.Format("Mon Jan _2 15:04:05 2006")))
}

func main() {
	fmt.Println("Server started .....")
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/ping", PingHandler)

	// Get port number from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3006"
	}

	fmt.Println("Listen serve for port... ", port)
	// in case of error, log.Fatal will exit application
	// Listen and Server in 0.0.0.0:3005
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
