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
    log.Println("start http listening :18888")
    httpServer.Addr = ":18888"
    log.Println(httpServer.ListenAndServe())
}
