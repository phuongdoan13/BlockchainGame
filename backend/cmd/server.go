package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/phuongdoan13/BlockchainGame/backend/api"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server is serving at http://localhost:8080/")
	// Init the mux router
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/matchingABI", api.MatchingContractABI)
	router.HandleFunc("/winERC20", api.WinERC20_controller).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
