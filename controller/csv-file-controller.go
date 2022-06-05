package controller

import (
	"encoding/csv"
	"encoding/json"
	"net/http"

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
	file, fileHeader, err := req.FormFile("csv")

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

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
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
