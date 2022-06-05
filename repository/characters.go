package repository

import (
	"encoding/json"
	"fmt"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/go-resty/resty/v2"
)

type CharacterClientRepository interface {
	FindCharacters() (*entity.ResponseBody, error)
}

type repo struct{}

func NewCharacterClientRepository() CharacterClientRepository {
	return &repo{}
}

func (c *repo) FindCharacters() (*entity.ResponseBody, error) {
	client := resty.New()
	resp, err := client.R().Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return nil, ErrorConnectingApi
	}

	responseBody := entity.ResponseBody{}
	err = json.Unmarshal([]byte(resp.Body()), &responseBody)
	if err != nil {
		fmt.Println(err.Error())
		return nil, ErrorConvertingToJSON
	}

	return &responseBody, nil
}
