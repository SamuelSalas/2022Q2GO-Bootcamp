package service

/*
import (
	"fmt"
	"testing"

	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCsvService_WorkerPoolReadCsv(t *testing.T) {
	tests := []struct {
		testName         string
		testData         [][]string
		idType           string
		items            int
		itemsWorkerLimit int
		errorResponse    error
	}{
		{
			testName:         "Only Valid Items",
			testData:         DataCharacters,
			idType:           "odd",
			items:            10,
			itemsWorkerLimit: 10,
			errorResponse:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			csvServiceImpl := NewCsvService(nil)
			results, err := csvServiceImpl.ReadCsvWorkerPool(test.idType, test.items, test.itemsWorkerLimit)
			fmt.Println(results.Results)
			assert.Equal(t, err, test.errorResponse)
		})
	}
}
*/
