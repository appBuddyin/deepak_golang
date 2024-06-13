package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Request struct {
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

var logger *log.Logger

func SetLogger(l *log.Logger) {
	logger = l
}

func logRequest(operation string, req Request, res Response) {
	entry := fmt.Sprintf(
		"Time: %s, Operation: %s, Operands: [%f, %f], Result: %f, Error: %s\n",
		time.Now().Format(time.RFC3339),
		operation,
		req.Operand1,
		req.Operand2,
		res.Result,
		res.Error,
	)
	logger.Println(entry)
}

func decodeRequest(w http.ResponseWriter, r *http.Request) (Request, bool) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logRequest("", req, Response{Error: "Invalid request body"})
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return req, false
	}
	return req, true
}
