package main

import (
	"grass/jsonapi/model"
	"log"
	"net/http"
)

func main() {

	router := model.NewRouter()
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
