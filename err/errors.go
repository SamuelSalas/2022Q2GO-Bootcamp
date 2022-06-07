package err

import "errors"

var (
	ErrorCsvInvalidColumnNumber = errors.New("invalid column number")
	ErrorCsvCreation            = errors.New("error creating csv")
	ErrorCsvReader              = errors.New("error reading csv")
	ErrorCsvEmpty               = errors.New("empty csv")
	ErrorConnectingApi          = errors.New("error connecting api")
	ErrorConvertingToJSON       = errors.New("error converting results to json")
	ErrorFileWasNotFound        = errors.New("file was not found")
	ErrorInvalidFileType        = errors.New("invalid file type")
	ErrorParameterNotFound      = errors.New("parameter not found")
	ErrorInvalidValueType       = errors.New("invalid value type")
	ErrorInvalidIdType          = errors.New("invalid id type")
)
