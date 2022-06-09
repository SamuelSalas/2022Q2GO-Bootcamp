package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
)

func (c *controller) GetCSVFileDataWorkerPool(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	workersParam, ok := req.URL.Query()["workers"]
	if !ok || len(workersParam[0]) < 1 {
		log.Println(err.ErrorParameterNotFound)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorParameterNotFound.Error()})
		return
	}

	workers, errs := strconv.Atoi(workersParam[0])
	if errs != nil {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorInvalidValueType.Error()})
		return
	}

	itemsParam, ok := req.URL.Query()["items"]
	if !ok || len(itemsParam[0]) < 1 {
		log.Println(err.ErrorParameterNotFound)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorParameterNotFound.Error()})
		return
	}

	items, errs := strconv.Atoi(itemsParam[0])
	if errs != nil {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorInvalidValueType.Error()})
		return
	}

	itemsPerWorkerParam, ok := req.URL.Query()["items_per_worker"]
	if !ok || len(itemsPerWorkerParam[0]) < 1 {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorParameterNotFound.Error()})
		return
	}

	itemsPerWorker, errs := strconv.Atoi(itemsPerWorkerParam[0])
	if errs != nil {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorInvalidValueType.Error()})
		return
	}

	idTypeParam, ok := req.URL.Query()["id_type"]
	if !ok || len(idTypeParam[0]) < 1 {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.ErrorParameterNotFound.Error()})
		return
	}

	result, errs := c.csvService.ReadCsvWorkerPool(idTypeParam[0], workers, items, itemsPerWorker)
	if errs != nil {
		log.Println(errs)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: errs.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result.Results)
}
