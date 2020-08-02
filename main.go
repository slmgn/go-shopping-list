package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping-list/middleware"
	"shopping-list/router"
)

func main() {

	middleware.CreateConnection()

	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
