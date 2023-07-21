package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mdNoman21/Go/Beginner-Projects/Task-Manager/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server at port 8080 .....")
	log.Fatal(http.ListenAndServe(":8080", r))

}
