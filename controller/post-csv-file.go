package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/utils"
)

func (*controller) PostCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	data, err := ValidateFile(req.FormFile("csv"))

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	result, err := csvService.ReadCsvData(data)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result.Results)
}
