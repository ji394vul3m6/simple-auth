package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Println("Start server")
	log.Fatal(http.ListenAndServe(":11180", router))
}
