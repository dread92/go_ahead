package goahead

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/ip", handler)
}
func handler(w http.ResponseWriter, r *http.Request) {
	header := r.RemoteAddr
	fmt.Fprint(w, header)
}

