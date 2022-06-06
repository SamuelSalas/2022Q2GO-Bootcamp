package service

import (
	"testing"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCsvService_ReadCsvFile(t *testing.T) {
	tests := []struct {
		testName       string
		expectedResult *entity.ResponseBody
		testData       [][]string
		errorResponse  error
	}{
		{
			testName:       "Happy Path",
			expectedResult: &ExpectedResult,
			testData:       DataCharacters,
			errorResponse:  nil,
		},
		{
			testName:       "Empty File",
			expectedResult: nil,
			testData:       [][]string{},
			errorResponse:  repository.ErrorCsvEmpty,
		},
		{
			testName:       "Invalid Column",
			expectedResult: nil,
			testData:       append(DataCharacters, []string{"extra info"}),
			errorResponse:  repository.ErrorCsvInvalidColumnNumber,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			csvServiceImpl := NewCsvService(nil)
			result, err := csvServiceImpl.ReadCsvData(test.testData)
			assert.Equal(t, *&test.expectedResult, result)
			assert.Equal(t, err, test.errorResponse)
		})
	}
}
