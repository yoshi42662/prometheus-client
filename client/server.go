package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieved to /")

	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
