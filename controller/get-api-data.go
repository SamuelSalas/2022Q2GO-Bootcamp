package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

func (c *controller) GetApiDataCsv(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := c.csvService.GenerateCsv()
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(entity.ErrorMessage{Message: err.Error()})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result.Results)
}
