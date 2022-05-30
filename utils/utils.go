package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func ConvertToJson(data [][]string) (character []entity.Character, err error) {
	for _, line := range data {
		columns := len(line)
		if columns != 7 {
			return nil, repository.ErrorCsvInvalidColumnNumber
		}

		var rec entity.Character = entity.Character{}
		rec.ID, err = strconv.Atoi(line[0])
		if err != nil {
			return nil, repository.ErrorCsvIDInvalidType
		}

		rec.Name = line[1]
		rec.Status = line[2]
		rec.Gender = line[3]
		rec.Image = line[4]
		rec.Url = line[5]
		rec.Created = line[6]

		character = append(character, rec)
	}
	return character, nil
}

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

func CreateTempFile(data [][]string) string {
	file, _ := os.Create("test.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, row := range data {
		_ = writer.Write(row)
	}
	return file.Name()
}
