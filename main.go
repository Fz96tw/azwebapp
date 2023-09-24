package main

import (
	"fmt"
	"net/http"
	"time"
)

var t = time.Now()

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/version", VersionServer)
	http.HandleFunc("/coder", CoderServer)
	http.ListenAndServe(":80", nil)
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func VersionServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "service version %s", t)
}

func CoderServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Coded by NH")
}
