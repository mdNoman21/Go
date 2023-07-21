// routes first then app.go, then utils,main.go,models,controller at last
package routes

import (
	"github.com/gorilla/mux"
	"github.com/mdNoman21/Go/Beginner-Projects/Bookstore-Management-API/pkg/controllers"
)

// func tha will have all routes ,routes which would help to reach controllers
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")

}
