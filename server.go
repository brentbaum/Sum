// httpserver.go
package main

import (
    "flag"
    "net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/parse", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })
    panic(http.ListenAndServe(":7331", nil))
}