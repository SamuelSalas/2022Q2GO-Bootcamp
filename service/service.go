package service

import (
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

type CsvService interface {
	ReadCsvData(data [][]string) (*entity.ResponseBody, error)
	GenerateCsv() error
	ReadCsvWorkerPool(data [][]string, idType string, items, items_per_workers int) (*entity.ResponseBody, error)
}
type csvService struct {
	repo repository.CharacterClientRepository
}

func NewCsvService(repository repository.CharacterClientRepository) CsvService {
	return &csvService{repository}
}
