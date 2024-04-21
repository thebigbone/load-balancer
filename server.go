package main

import (
    "fmt"
    "os"
    "flag"
    "log"
    "net/http"
    "net/http/httputil"
)

func main() {
    port := flag.Int("port", 8081, "Port to listen on")
    flag.Parse()

    fmt.Println("hello world!") 
    http.HandleFunc("/hello", hello)

    fmt.Printf("starting the server at port %d\n", *port)
    if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
        log.Fatal(err)
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    r.Header.Add("test-header", "test-value")

    logger := log.New(os.Stdout, "", 0)

    reqDump, err := httputil.DumpRequest(r, true)

    if err != nil {
        log.Fatal(err)
    }
    
    logger.Printf("Received request from: %s\n", r.RemoteAddr)
    logger.Printf("\n%s", string(reqDump))
    fmt.Printf("hello world!")

}
