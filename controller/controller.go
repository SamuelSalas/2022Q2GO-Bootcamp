package controller

import (
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
)

type CSVController interface {
	PostCSVFile(resp http.ResponseWriter, req *http.Request)
	GetApiDataCsv(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var csvService service.CsvService

func NewCsvController(service service.CsvService) CSVController {
	csvService = service
	return &controller{}
}
