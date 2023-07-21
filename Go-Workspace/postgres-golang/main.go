package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdNoman21/Go/Beginner-Projects/postgres-golang/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at the port 8080.....")

	log.Fatal(http.ListenAndServe(":8080", r))
}
