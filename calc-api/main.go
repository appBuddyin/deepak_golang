package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Request struct {
	Operation string  `json:"operation"`
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

var logger *log.Logger

func main() {
	// Initialize logger
	file, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger = log.New(file, "", log.LstdFlags)

	r := mux.NewRouter()
	r.HandleFunc("/calculate", calculateHandler).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logRequest(req, Response{Error: "Invalid request body"})
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var res Response
	switch req.Operation {
	case "add":
		res.Result = req.Operand1 + req.Operand2
	case "subtract":
		res.Result = req.Operand1 - req.Operand2
	case "multiply":
		res.Result = req.Operand1 * req.Operand2
	case "divide":
		if req.Operand2 == 0 {
			res.Error = "Division by zero"
		} else {
			res.Result = req.Operand1 / req.Operand2
		}
	default:
		res.Error = "Invalid operation"
	}

	logRequest(req, res)

	if res.Error != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func logRequest(req Request, res Response) {
	entry := fmt.Sprintf(
		"Time: %s, Request: %s, Operands: [%f, %f], Response: %f, Error: %s\n",
		time.Now().Format(time.RFC3339),
		req.Operation,
		req.Operand1,
		req.Operand2,
		res.Result,
		res.Error,
	)
	logger.Println(entry)
}
