package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "It's working, now go write some actual code you lazy person you.")
    l := log.New(os.Stdout, "[AAA] ", log.Ldate | log.Ltime)
    l.Printf("Request")
}

# Another comment

func main() {
    l := log.New(os.Stdout, "[AAA] ", log.Ldate | log.Ltime)
    l.Printf("Listening on localhost")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
