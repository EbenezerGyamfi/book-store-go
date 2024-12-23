package routes

import (
	"example.com/bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)


func RegisterBook(router *mux.Router)  {
	

	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}