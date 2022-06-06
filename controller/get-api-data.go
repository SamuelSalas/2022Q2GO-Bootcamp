package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

func (*controller) GetApiDataCsv(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	err := csvService.GenerateCsv()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(entity.Message{Message: "succeed"})
}
