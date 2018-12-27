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
	// 静态文件路由
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	// 普通路由
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "this is home")
}
