package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	port := "8080"
	log.Println("Starting ballerina action on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
