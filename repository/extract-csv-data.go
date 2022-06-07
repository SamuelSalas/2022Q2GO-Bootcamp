package repository

import (
	"encoding/csv"
	"os"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
)

func (c *repo) ExtractCsvData() (*[][]string, error) {
	file, errs := os.Open("result.csv")
	if errs != nil {
		return nil, err.ErrorFileWasNotFound
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, errs := csvReader.ReadAll()
	if errs != nil {
		return nil, err.ErrorCsvReader
	}
	if len(data) == 0 {
		return nil, err.ErrorCsvEmpty
	}

	return &data, nil
}
