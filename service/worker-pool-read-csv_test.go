package service

import (
	"testing"

	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCsvService_WorkerPoolReadCsv(t *testing.T) {
	tests := []struct {
		testName         string
		testData         *[][]string
		idType           string
		items            int
		itemsWorkerLimit int
		errorResponse    error
	}{
		{
			testName:         "Only Valid Items",
			testData:         &DataCharacters,
			idType:           "odd",
			items:            20,
			itemsWorkerLimit: 2,
			errorResponse:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			csvServiceMockImpl := CsvServiceMock{}
			csvServiceMockImpl.On("ExtractCsvData").Return(test.testData, test.errorResponse)
			csvServiceImpl := NewCsvService(&csvServiceMockImpl)
			_, err := csvServiceImpl.ReadCsvWorkerPool(test.idType, test.items, test.itemsWorkerLimit)
			assert.Equal(t, err, test.errorResponse)
		})
	}
}
