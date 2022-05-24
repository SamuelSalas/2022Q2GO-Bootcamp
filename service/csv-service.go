package service

import (
	"encoding/json"
	"fmt"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/go-resty/resty/v2"
)

type CsvService interface {
	ConvertCsvToJson(data [][]string) ([]*entity.CSV, error)
	RequestRickAndMortyCharacters() (*entity.ResponseBody, error)
}
type service struct{}

func NewCsvService() CsvService {
	return &service{}
}

func (*service) ConvertCsvToJson(data [][]string) ([]*entity.CSV, error) {
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

func (*service) RequestRickAndMortyCharacters() (*entity.ResponseBody, error) {
	client := resty.New()
	resp, err := client.R().Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	responseBody := entity.ResponseBody{}
	err = json.Unmarshal(resp.Body(), &responseBody)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &responseBody, nil
}
