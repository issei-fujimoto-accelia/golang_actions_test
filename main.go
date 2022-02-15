package main

import(
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func main() {
    var httpServer http.Server
    http.HandleFunc("/", handler)
    log.Println("start http listening :8080")
    httpServer.Addr = ":8080"
    log.Println(httpServer.ListenAndServe())
}
