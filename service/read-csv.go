package service

import (
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func (*csvService) ReadCsvData(data [][]string) (*entity.ResponseBody, error) {
	responseBody := entity.ResponseBody{}
	if len(data) == 0 {
		return nil, repository.ErrorCsvEmpty
	}

	for _, line := range data {
		if len(line) != 7 {
			return nil, repository.ErrorCsvInvalidColumnNumber
		}

		var rec entity.Character = entity.Character{}
		rec.ID, _ = strconv.Atoi(line[0])
		rec.Name = line[1]
		rec.Status = line[2]
		rec.Gender = line[3]
		rec.Image = line[4]
		rec.Url = line[5]
		rec.Created = line[6]
		responseBody.Results = append(responseBody.Results, rec)
	}
	return &responseBody, nil
}
