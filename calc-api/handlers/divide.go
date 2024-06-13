package handlers

import (
	"encoding/json"
	"net/http"
	"calc-api/Service"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}

	service := Service.CreateOperationInstance()
	result, err := service.Divide(req.Operand1, req.Operand2)
	var res Response
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Result = result
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
