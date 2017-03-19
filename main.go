package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

// var re = regexp.MustCompile("^/(.*)@golang.org$")

//GoperSuffix - const which represent part for gopher detection
const GoperSuffix = "@golang.org"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	path := r.URL.Path[1:]
	if strings.HasSuffix(path, GoperSuffix) {
		fmt.Fprintf(w, "Hello, gopher %s!\n", strings.TrimSuffix(path, GoperSuffix))
	}
	fmt.Fprintf(w, "Hello, %s!\n", path)
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
