package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/", fileServer)
	PORT := ":8000"
	fmt.Printf("Server started on port %v", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
