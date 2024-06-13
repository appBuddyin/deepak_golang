package handlers

import (
	"calc-api/Service"
	"encoding/json"
	"net/http"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}

	operation := Service.CreateOperationInstance()
	res := Response{
		Result: operation.Subtract(req.Operand1, req.Operand2),
	}

	logRequest("subtract", req, res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
