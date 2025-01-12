package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	fmt.Println("hello")
	http.ListenAndServe(":8080", nil)
}
