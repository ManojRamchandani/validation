package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	httpAddr := "0.0.0.0:9999"

	port := os.Getenv("PORT")
	if port != "" {
		httpAddr = "0.0.0.0:" + port
	}
	fmt.Println("Http Server is listening on", httpAddr)

	// Start the http server using the custom router
	exitMsg := http.ListenAndServe(httpAddr, setupRouter())
	log.Fatal(exitMsg)

}

func setupRouter() *mux.Router {

	// Create a new empty http router
	httpRouter := mux.NewRouter()

	// When an http GET request is received on the path /health, delegate to function "HealthCheckHandler()"
	httpRouter.HandleFunc("/health", HealthCheckHandler)
	//httpRouter.HandleFunc("/", HomeHandler)

	return httpRouter

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Print a line in the logs
	fmt.Println("HIT: healthcheck")

	// Write the string "ALIVE" into the responses body
	_, _ = io.WriteString(w, "ALIVE")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Print a line in the logs
	fmt.Println("HIT: homepage")

	// Write the string "ALIVE" into the responses body
	_, _ = io.WriteString(w, "THIS IS THE DEFAULT HOMEPAGE")
}
