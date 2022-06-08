package service

import (
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type CsvService interface {
	ReadCsvData() (*entity.ResponseBody, error)
	GenerateCsv() (*entity.ResponseBody, error)
	ReadCsvWorkerPool(idType string, items, items_per_workers int) (*entity.ResponseBody, error)
}

type csvService struct {
	csvRepo csvRepository
}

type csvRepository interface {
	FindCharacters() (*entity.ResponseBody, error)
	ExtractCsvData() (*[][]string, error)
}

func NewCsvService(csvRepo csvRepository) CsvService {
	return &csvService{
		csvRepo,
	}
}
