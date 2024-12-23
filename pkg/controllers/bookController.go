package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/bookstore/pkg/models"
	"example.com/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)



func CreateBook(res http.ResponseWriter, req *http.Request)  {

	createBook := &models.Book{}
	utils.ParseBody(req, createBook)

	b := createBook.Save()

	jsonData, err := json.Marshal(b)	

	if err!= nil {
        http.Error(res, "Failed to marshal book data", http.StatusInternalServerError)
        return
    }

	res.Header().Set("Content-Type", "application/json")

	res.Write(jsonData)
}


func GetBooks(res http.ResponseWriter, req *http.Request) {
    log.Println("GetBooks endpoint hit")

    // Retrieve the books from the model
    newBooks, err := models.GetBooks()
    if err != nil {
        log.Printf("Error fetching books: %v\n", err)
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }

    // Marshal the books into JSON
    jsonData, err := json.Marshal(newBooks)
    if err != nil {
        log.Printf("Error marshalling books: %v\n", err)
        http.Error(res, "Failed to marshal books data", http.StatusInternalServerError)
        return
    }

    // Write the response
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    res.Write(jsonData)
}





func GetBookById(res http.ResponseWriter, req *http.Request)  {

	vars := mux.Vars(req)
	bookId := vars["id"]

	
	ID, err := strconv.ParseInt(bookId, 10, 64)

	if err!= nil {
        http.Error(res, "Invalid book ID", http.StatusBadRequest)
        return
    }

	bookDetails, _ := models.GetBookById(ID)

	if bookDetails == nil {
        http.Error(res, "Book not found", http.StatusNotFound)
        return
    }

	jsonData, err := json.Marshal(bookDetails)	

	if err!= nil {
        http.Error(res, "Failed to marshal book data", http.StatusInternalServerError)
        return
    }

	res.Header().Set("Content-Type", "application/json")

	res.Write(jsonData)
    
}



func UpdateBook(res http.ResponseWriter, req *http.Request)  {

	vars := mux.Vars(req)
	bookId := vars["id"]


    ID, err := strconv.ParseInt(bookId, 10, 64)

	if err!= nil {
        http.Error(res, "Invalid book ID", http.StatusBadRequest)
        return
    }

	updateBook := &models.Book{}

	utils.ParseBody(req, updateBook)

	updateBook.ID = uint(ID)

	bookDetails := updateBook.Update(int64(updateBook.ID))

	jsonData, err := json.Marshal(bookDetails)

	if err!= nil {
        http.Error(res, "Failed to marshal book data", http.StatusInternalServerError)
        return
    }

	res.Header().Set("Content-Type", "application/json")

	res.Write(jsonData)
    
}


func DeleteBook(res http.ResponseWriter, req *http.Request)  {

	vars := mux.Vars(req)

	bookId := vars["id"]


    ID, err := strconv.ParseInt(bookId, 10, 64)

	if err!= nil {
        http.Error(res, "Invalid book ID", http.StatusBadRequest)
        return
    }

	bookDetails := models.DeleteBook(ID)

	jsonData, err := json.Marshal(bookDetails)

	if err!= nil {
        http.Error(res, "Failed to marshal book data", http.StatusInternalServerError)
        return
    }

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)

	res.WriteHeader(http.StatusNoContent)
    
}