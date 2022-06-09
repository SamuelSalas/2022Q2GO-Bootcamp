package service

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
)

func (s *csvService) GenerateCsv() (*entity.ResponseBody, error) {
	result, errs := s.csvRepo.FindCharacters()
	if errs != nil {
		return nil, errs
	}

	errs = structToCsv(result.Results)
	if errs != nil {
		return nil, errs
	}

	return result, nil
}

func structToCsv(characters []entity.Character) error {
	file, errs := os.Create("result.csv")
	if errs != nil {
		return err.ErrorCsvCreation
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, character := range characters {
		var row []string
		row = append(row, strconv.Itoa(character.ID))
		row = append(row, character.Name)
		row = append(row, character.Status)
		row = append(row, character.Gender)
		row = append(row, character.Image)
		row = append(row, character.Url)
		row = append(row, character.Created)
		writer.Write(row)
	}
	return nil
}
