package handlers

import (
	"encoding/json"
	"net/http"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	
	var res Response
	if req.Operand2 == 0 {
		res.Error = "Division by zero"
	} else {
		res.Result = req.Operand1 / req.Operand2
	}

	logRequest("divide", req, res)
	if res.Error != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(res)
}
