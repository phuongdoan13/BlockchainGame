package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/phuongdoan13/BlockchainGame/backend/api"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main(){
	fmt.Println("Server is serving at http://localhost:8080/")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/matchingABI", api.MatchingContractABI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}