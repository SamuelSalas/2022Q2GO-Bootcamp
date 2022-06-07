package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
)

func (*controller) PostWorkerPoolCSVFile(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	data, err := validateFile(req.FormFile("csv"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	itemsParam, ok := req.URL.Query()["items"]
	if !ok || len(itemsParam[0]) < 1 {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorParameterNotFound.Error()})
		return
	}

	items, err := strconv.Atoi(itemsParam[0])
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorInvalidValueType.Error()})
		return
	}

	itemsPerWorkerParam, ok := req.URL.Query()["items_per_worker"]
	if !ok || len(itemsPerWorkerParam[0]) < 1 {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorParameterNotFound.Error()})
		return
	}

	itemsPerWorker, err := strconv.Atoi(itemsPerWorkerParam[0])
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorInvalidValueType.Error()})
		return
	}

	idTypeParam, ok := req.URL.Query()["id_type"]
	if !ok || len(idTypeParam[0]) < 1 {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorParameterNotFound.Error()})
		return
	}

	if idTypeParam[0] != "odd" && idTypeParam[0] != "even" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: repository.ErrorInvalidIdType.Error()})
		return
	}

	result, err := csvService.ReadCsvWorkerPool(data, idTypeParam[0], items, itemsPerWorker)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result.Results)
}
