package handlers

import (
	"encoding/json"
	"net/http"
	"calc-api/Service"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}

	service := Service.CreateOperationInstance()
	res := Response{
		Result: service.Subtract(req.Operand1, req.Operand2),
	}

	logRequest("subtract", req, res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
