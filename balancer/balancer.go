package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

var servers = []string{
    "127.0.0.1:8081",
    "127.0.0.1:8082",
}

var counter uint64

func main() {
    http.HandleFunc("/", loadBalancer)
    log.Println("starting load balancer at 8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func loadBalancer(w http.ResponseWriter, r *http.Request) {
    // Select the next server to handle the request
    index := counter % uint64(len(servers))
    serverURL := servers[index]
    counter++
    // Create a reverse proxy to forward the request to the selected server
    proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: serverURL})
    proxy.ServeHTTP(w, r)

    // Log the request and the selected server
    fmt.Printf("Received request from %s, forwarding to %s\n", r.RemoteAddr, serverURL)
}

