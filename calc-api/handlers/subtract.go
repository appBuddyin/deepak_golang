package handlers

import (
	"encoding/json"
	"net/http"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	
	res := Response{
		Result: req.Operand1 - req.Operand2,
	}

	logRequest("subtract", req, res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
