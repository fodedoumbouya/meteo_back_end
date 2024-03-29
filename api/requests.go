package api

import (
	"encoding/json"
	"fmt"
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

// /widget?id
func GetWidget(w http.ResponseWriter, r *http.Request) {
	//id, ok := r.URL.Query()["id"]

	id := r.URL.Query().Get("id")
	html := "This request need id"
	//r.URL.Host
	if len(id) > 0 {
		html = getHtml(id)

	}
	//"error"
	// fmt.Println(html)

	fmt.Fprintf(w, html)

}
