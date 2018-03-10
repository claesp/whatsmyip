package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")])
	})

	http.ListenAndServe(":8081", nil)
}
