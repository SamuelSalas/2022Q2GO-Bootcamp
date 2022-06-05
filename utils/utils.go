package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func GenerateCsv(characters *[]entity.Character) error {
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

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
