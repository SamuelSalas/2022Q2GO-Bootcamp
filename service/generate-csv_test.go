package service

import (
	"os"
	"testing"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type csvServiceMock struct {
	mock.Mock
}

func (m *csvServiceMock) FindCharacters() (*entity.ResponseBody, error) {
	args := m.Called()
	return args.Get(0).(*entity.ResponseBody), args.Error(1)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

var fileName = "result.csv"

func TestCsvService_RequestRickAndMortyCharacters_TestTable(t *testing.T) {
	tests := []struct {
		testName       string
		expectedResult *entity.ResponseBody
		errorResponse  error
		fileExists     bool
	}{
		{
			testName:       "Happy Path",
			expectedResult: &ExpectedResult,
			errorResponse:  nil,
			fileExists:     true,
		},
		{
			testName:       "Unable to Connect to API",
			expectedResult: nil,
			errorResponse:  repository.ErrorConnectingApi,
			fileExists:     false,
		},
	}
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			defer os.Remove(fileName)
			csvServiceMockImpl := csvServiceMock{}
			csvServiceMockImpl.On("FindCharacters").Return(test.expectedResult, test.errorResponse)
			csvServiceImpl := NewCsvService(&csvServiceMockImpl)
			err := csvServiceImpl.GenerateCsv()
			isFileExists := fileExists(fileName)
			assert.Equal(t, err, test.errorResponse)
			assert.Equal(t, isFileExists, test.fileExists)
		})
	}
}
