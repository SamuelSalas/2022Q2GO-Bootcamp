package controller

import (
	"encoding/json"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/errors"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
)

type CSVFileController interface {
	PostCSVFile(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var csvService service.CsvService

func NewCsvController(service service.CsvService) CSVFileController {
	csvService = service
	return &controller{}
}

func (*controller) PostCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	file, fileHeader, fileError := req.FormFile("csv")

	if fileError != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: fileError.Error()})
		return
	}

	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: "Invalid file type"})
		return
	}

	result, err := csvService.ReadCsvFile(file)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ErrorMessage{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
