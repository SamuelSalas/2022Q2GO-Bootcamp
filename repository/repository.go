package repository

import (
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type CsvRepository interface {
	FindCharacters() (*entity.ResponseBody, error)
	ExtractCsvData() (*[][]string, error)
}

type repo struct{}

func NewCsvRepository() CsvRepository {
	return &repo{}
}
