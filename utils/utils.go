package utils

import (
	"fmt"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
)

func ConvertToJson(data [][]string) ([]*entity.CSV, error) {
	var csvData []*entity.CSV
	for _, line := range data {
		columns := len(line)
		if columns != 2 {
			return nil, fmt.Errorf("invalid column number: %d", columns)
		}

		var rec *entity.CSV = &entity.CSV{
			ID:    line[0],
			Items: line[1],
		}

		csvData = append(csvData, rec)
	}
	return csvData, nil
}
