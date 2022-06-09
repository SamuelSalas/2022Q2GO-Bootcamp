package controller

import (
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

type CSVController interface {
	GetCSVFileData(resp http.ResponseWriter, req *http.Request)
	GetApiDataCsv(resp http.ResponseWriter, req *http.Request)
	GetCSVFileDataWorkerPool(resp http.ResponseWriter, req *http.Request)
}

type controller struct {
	csvService serviceCsv
}

type serviceCsv interface {
	ReadCsvData() (*entity.ResponseBody, error)
	GenerateCsv() (*entity.ResponseBody, error)
	ReadCsvWorkerPool(idType string, workers, items, items_per_workers int) (*entity.ResponseBody, error)
}

func NewCsvController(service serviceCsv) CSVController {
	return &controller{
		service,
	}
}
