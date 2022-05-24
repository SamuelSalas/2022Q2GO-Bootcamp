package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type CsvService interface {
	ConvertCsvToJson(file multipart.File) ([]*entity.CSV, error)
}
type service struct{}

func NewCsvService() CsvService {
	return &service{}
}

func (*service) ConvertCsvToJson(file multipart.File) ([]*entity.CSV, error) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	var csvData []*entity.CSV

	for _, line := range data {
		columns := len(line)
		if columns != 2 {
			return nil, fmt.Errorf("invalid column number: %d", columns)
		}

		var rec *entity.CSV = &entity.CSV{
			ID:    line[0],
			Items: line[1],
		}

		csvData = append(csvData, rec)
	}
	return csvData, nil
}
