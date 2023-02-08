package main

import (
	"log"
	"net/http"

	"github.com/deyanEnchev/src/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleJobs)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
