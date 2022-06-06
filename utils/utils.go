package utils

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func ValidateFile(file multipart.File, fileHeader *multipart.FileHeader, err error) ([][]string, error) {
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
