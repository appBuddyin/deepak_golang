package main

import (
	"log"
	"net/http"
	"os"

	"calc-api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	file, err := os.OpenFile("log/requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	handlers.SetLogger(logger)

	r := mux.NewRouter()
	r.HandleFunc("/add", handlers.AddHandler).Methods("POST")
	r.HandleFunc("/subtract", handlers.SubtractHandler).Methods("POST")
	r.HandleFunc("/multiply", handlers.MultiplyHandler).Methods("POST")
	r.HandleFunc("/divide", handlers.DivideHandler).Methods("POST")

	log.Println("Server started on :3333")
	log.Fatal(http.ListenAndServe(":3333", r))
}
