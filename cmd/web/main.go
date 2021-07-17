package main

import (
	"log"
	"net/http"
)

func main() {
	routes := routes()

	log.Println("Starting a web server on port :8000")
	_ = http.ListenAndServe(":8000", routes)
}
