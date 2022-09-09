package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{vin}", crawler.engineCrawler)

	fmt.Println("Running... localhost:8080/{vin}")
	http.ListenAndServe(":8080", router)
}
