package service

import (
	"os"
	"testing"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/entity"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/test/testdata"
	. "github.com/SamuelSalas/2022Q2GO-Bootcamp/utils"
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

func TestCsvService_ReadCsvFile(t *testing.T) {
	tmpFile, _ := os.Open(CreateTempFile(TestCharacters))
	os.Remove(tmpFile.Name())
	dataExpected, _ := ConvertToJson(TestCharacters)
	csvServiceImpl := NewCsvService(nil)
	result, err := csvServiceImpl.ReadCsvFile(tmpFile)
	assert.Equal(t, &dataExpected, result)
	assert.Nil(t, err)
}

func TestCsvService_ReadCsvFile_EmptyFile(t *testing.T) {
	tmpFile, _ := os.Open(CreateTempFile(nil))
	os.Remove(tmpFile.Name())
	csvServiceImpl := NewCsvService(nil)
	result, err := csvServiceImpl.ReadCsvFile(tmpFile)
	assert.Empty(t, result)
	assert.EqualError(t, err, repository.ErrorCsvEmpty.Error())
}

func TestCsvService_ReadCsvFile_InvalidColumnNumber(t *testing.T) {
	testData := TestCharacters
	for i, character := range testData {
		testData[i] = append(character, "extra info")
	}

	tmpFile, _ := os.Open(CreateTempFile(TestCharacters))
	os.Remove(tmpFile.Name())
	csvServiceImpl := NewCsvService(nil)
	result, err := csvServiceImpl.ReadCsvFile(tmpFile)
	assert.Empty(t, result)
	assert.EqualError(t, err, repository.ErrorCsvInvalidColumnNumber.Error())
}

func TestCsvService_ReadCsvFile_InvalidField(t *testing.T) {
	d := len(TestCharacters)
	testData := TestCharacters
	testData[d-1] = append(testData[d-1], "s")
	tmpFile, _ := os.Open(CreateTempFile(TestCharacters))
	os.Remove(tmpFile.Name())
	csvServiceImpl := NewCsvService(nil)
	result, err := csvServiceImpl.ReadCsvFile(tmpFile)
	assert.Empty(t, result)
	assert.EqualError(t, err, repository.ErrorCsvReader.Error())
}

func TestCsvService_RequestRickAndMortyCharacters(t *testing.T) {
	defer os.Remove("result.csv")
	csvServiceMockImpl := csvServiceMock{}
	characters, _ := ConvertToJson(TestCharacters)
	dataExpected := entity.ResponseBody{
		Results: characters,
		Info:    InfoData,
	}
	csvServiceMockImpl.On("FindCharacters").Return(&dataExpected, nil)
	csvServiceImpl := NewCsvService(&csvServiceMockImpl)
	err := csvServiceImpl.RequestRickAndMortyCharacters()
	assert.Nil(t, err)
}
