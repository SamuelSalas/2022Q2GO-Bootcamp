package service

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func (c *csvService) GenerateCsv() error {
	result, err := c.repo.FindCharacters()
	if err != nil {
		return err
	}

	err = structToCsv(&result.Results)
	if err != nil {
		return err
	}

	return nil
}

func structToCsv(characters *[]entity.Character) error {
	file, err := os.Create("result.csv")
	if err != nil {
		return repository.ErrorCsvCreation
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, character := range *characters {
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
