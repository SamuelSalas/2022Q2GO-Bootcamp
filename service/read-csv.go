package service

import (
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
)

func (s *csvService) ReadCsvData() (*entity.ResponseBody, error) {
	responseBody := entity.ResponseBody{}
	data, errs := s.csvRepo.ExtractCsvData()
	if errs != nil {
		return nil, errs
	}

	character := entity.Character{}
	for _, row := range *data {
		if len(row) != 7 {
			return nil, err.ErrorCsvInvalidColumnNumber
		}

		character.ID, _ = strconv.Atoi(row[0])
		character.Name = row[1]
		character.Status = row[2]
		character.Gender = row[3]
		character.Image = row[4]
		character.Url = row[5]
		character.Created = row[6]
		responseBody.Results = append(responseBody.Results, character)
	}

	return &responseBody, nil
}
