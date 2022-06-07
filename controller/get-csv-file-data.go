package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

func (c *controller) GetCSVFileData(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := c.csvService.ReadCsvData()
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.Message{Message: err.Error()})
		return
	}

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result.Results)
}
