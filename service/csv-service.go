package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/utils"
)

type CsvService interface {
	ReadCsvFile(file multipart.File) ([]*entity.CSV, error)
}
type service struct{}

func NewCsvService() CsvService {
	return &service{}
}

func (*service) ReadCsvFile(file multipart.File) ([]*entity.CSV, error) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	csvData, err := ConvertToJson(data)
	return csvData, nil
}
