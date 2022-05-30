package service

import (
	"encoding/csv"
	"os"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/utils"
)

type CsvService interface {
	ReadCsvFile(file *os.File) (*[]entity.Character, error)
	RequestRickAndMortyCharacters() error
}
type csvService struct {
	repo repository.CharacterClientRepository
}

func NewCsvService(repository repository.CharacterClientRepository) CsvService {
	return &csvService{repository}
}

func (*csvService) ReadCsvFile(file *os.File) (*[]entity.Character, error) {
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, repository.ErrorCsvReader
	}

	if len(data) == 0 {
		return nil, repository.ErrorCsvEmpty
	}

	csvData, err := ConvertToJson(data)
	if err != nil {
		return nil, err
	}

	return &csvData, nil
}

func (c *csvService) RequestRickAndMortyCharacters() error {
	result, err := c.repo.FindCharacters()
	if err != nil {
		return err
	}

	err = GenerateCsv(&result.Results)
	if err != nil {
		return err
	}

	return nil
}
