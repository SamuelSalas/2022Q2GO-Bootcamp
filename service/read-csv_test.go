package service

import (
	"testing"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/err"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	"github.com/stretchr/testify/assert"
)

func (m *CsvServiceMock) ExtractCsvData() (*[][]string, error) {
	args := m.Called()
	return args.Get(0).(*[][]string), args.Error(1)
}

var extraTestData [][]string = append(DataCharacters, []string{"extra info"})

func TestCsvService_ReadCsvFile(t *testing.T) {
	tests := []struct {
		testName       string
		expectedResult *entity.ResponseBody
		testData       *[][]string
		errorResponse  error
	}{
		{
			testName:       "Happy Path",
			expectedResult: &ExpectedResult,
			testData:       &DataCharacters,
			errorResponse:  nil,
		},
		{
			testName:       "Empty File",
			expectedResult: nil,
			testData:       nil,
			errorResponse:  err.ErrorCsvEmpty,
		},
		{
			testName:       "Invalid Column",
			expectedResult: nil,
			testData:       &extraTestData,
			errorResponse:  err.ErrorCsvInvalidColumnNumber,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			csvServiceMockImpl := CsvServiceMock{}
			csvServiceMockImpl.On("ExtractCsvData").Return(test.testData, test.errorResponse)
			csvServiceImpl := NewCsvService(&csvServiceMockImpl)
			result, errs := csvServiceImpl.ReadCsvData()
			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, errs, test.errorResponse)
		})
	}
}
