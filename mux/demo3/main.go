package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Run()
}

func Run() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	// 初始化Router
	r := mux.NewRouter()
	// 指定host
	r.HandleFunc("/host", HostHandler).Host("localhost")
	http.ListenAndServe(":3000", r)
}

func HostHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "the host is www.example.com")
}
