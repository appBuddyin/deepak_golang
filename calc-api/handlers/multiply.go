package handlers

import (
	"calc-api/Service"
	"encoding/json"
	"net/http"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}

	service := Service.CreateOperationInstance()
	res := Response{
		Result: service.Multiply(req.Operand1, req.Operand2),
	}

	logRequest("Multiply", req, res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
