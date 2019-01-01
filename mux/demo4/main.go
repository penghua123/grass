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
	// 带变量的url路由
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("Get").Name("user")

	url, _ := r.Get("user").URL("id", "5")
	fmt.Println("user url: ", url.String())
	http.ListenAndServe(":3000", r)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "this is get user, and the user id is ", vars["id"])
}
