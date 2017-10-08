package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/prometheus/client_golang/prometheus/promhttp"
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
