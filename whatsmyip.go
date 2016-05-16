package main

import (
	"fmt"
	"strings"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	var ip string
	rIp := r.Header["X-Real-Ip"]

	if len(rIp) == 0 {
		ip = r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")]
	} else {
		ip = rIp[0]
	}
	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/", root) 
	http.ListenAndServe(":30001", nil)
}
