package controller

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
)

type CSVController interface {
	PostCSVFile(resp http.ResponseWriter, req *http.Request)
	GetRickAndMortyCharactersCsv(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var csvService service.CsvService

func NewCsvController(service service.CsvService) CSVController {
	csvService = service
	return &controller{}
}

func (*controller) PostCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	_, fileHeader, err := req.FormFile("csv")

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: "Invalid file type"})
		return
	}

	file, err := os.Open(fileHeader.Filename)
	result, err := csvService.ReadCsvFile(file)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

func (*controller) GetRickAndMortyCharactersCsv(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	err := csvService.RequestRickAndMortyCharacters()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(entity.Message{Message: "succeed"})
}
