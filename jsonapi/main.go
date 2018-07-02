package main

import (
	"grass/jsonapi/model"
	"log"
	"net/http"
)

func main() {

	router := model.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
