package main

import (
	"fmt"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/router"
	"github.com/gorilla/mux"
)

const (
	port string = ":8080"
)

func main() {
	r := mux.NewRouter()
	router.Router(r)
	fmt.Printf("Mux HTTP Server running on port: %v", port)
	http.ListenAndServe(port, r)
}
