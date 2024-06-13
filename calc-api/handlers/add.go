package handlers

import (
	"calc-api/Service"
	"encoding/json"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}

	service := Service.CreateOperationInstance()
	res := Response{
		Result: service.Add(req.Operand1, req.Operand2),
	}

	logRequest("add", req, res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
