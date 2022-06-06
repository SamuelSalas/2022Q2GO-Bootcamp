package controller

import (
	"encoding/csv"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func (*controller) PostCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	data, err := validateFile(req.FormFile("csv"))

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

func validateFile(file multipart.File, fileHeader *multipart.FileHeader, err error) ([][]string, error) {
	data := [][]string{}
	if err != nil {
		return data, repository.ErrorFileWasNotFound
	}

	if fileHeader.Header.Get("Content-Type") != "text/csv" {
		return data, repository.ErrorInvalidFileType
	}

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()
	if err != nil {
		return data, repository.ErrorCsvReader
	}
	return data, nil
}
