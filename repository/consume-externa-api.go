package repository

import (
	"encoding/json"
	"fmt"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
	"github.com/go-resty/resty/v2"
)

func (c *repo) FindCharacters() (*entity.ResponseBody, error) {
	client := resty.New()
	resp, errs := client.R().Get("https://rickandmortyapi.com/api/character")
	if errs != nil {
		return nil, err.ErrorConnectingApi
	}

	responseBody := entity.ResponseBody{}
	errs = json.Unmarshal(resp.Body(), &responseBody)
	if errs != nil {
		fmt.Println(errs.Error())
		return nil, err.ErrorConvertingToJSON
	}

	return &responseBody, nil
}
