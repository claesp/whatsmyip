package main

import (
	"fmt"
	"strings"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]
	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/", root) 
	http.ListenAndServe(":30001", nil)
}
