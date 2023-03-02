package api

import (
	"encoding/json"
	"meteo_back_end/constant"
	"meteo_back_end/models"
	"net/http"
)

func GetStation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stations, hasError := getAllStationData()

	msg := "success"
	code := constant.OKstatus

	if hasError {
		msg = "error to get data"
		code = constant.ERRORstatus
	}

	resp := models.ResponseRquest{
		Message: msg,
		Code:    code,
		Data:    stations,
	}

	json.NewEncoder(w).Encode(resp)
}
