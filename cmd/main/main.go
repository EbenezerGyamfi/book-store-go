package main

import (
    "log"
    "net/http"

    "example.com/bookstore/pkg/routes"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    routes.RegisterBook(router)

    log.Println("Server started on port 8081")
    log.Fatal(http.ListenAndServe(":8081", router))
}
