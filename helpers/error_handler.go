package helpers

import (
	"encoding/json"
	"fmt"
	"go_api/proto_gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

type APIErrorMessage struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func loadErrorMessages() map[int]APIErrorMessage {

	filename := GetDirPath("error_messages.json")

	var apiErrorList []APIErrorMessage

	jsonFile, err := os.ReadFile(filename)
	ErrorHandler(err)

	err = json.Unmarshal(jsonFile, &apiErrorList)
	ErrorHandler(err)

	APIErrorMap := map[int]APIErrorMessage{}

	for i := 0; i < len(apiErrorList); i++ {
		apiError := apiErrorList[i]
		APIErrorMap[apiError.StatusCode] = apiError
	}

	return APIErrorMap
}

var APIErrorMap = loadErrorMessages()

func HttpCode2GrpcCode(code int) codes.Code {
	switch code {
	case 400:
		return 3
	case 401:
		return 16
	case 404:
		return 5
	case 403:
		return 7
	case 499:
		return 1
	case 504:
		return 4
	case 409:
		return 6
	case 501:
		return 11
	case 503:
		return 14
	default:
		return 13
	}
}

func Error(c int) error {
	apiError, ok := APIErrorMap[c]
	if !ok {
		println("Wrong Code")
	}

	errResponseFormat := proto_gen.ErrorResponse{
		Code:       int64(apiError.Code),
		StatusCode: int64(apiError.StatusCode),
		Message:    apiError.Message,
	}

	errorStatus, err := status.New(HttpCode2GrpcCode(apiError.Code), apiError.Message).WithDetails(&errResponseFormat)

	if err != nil {
		println("Something Wrong")
	}

	return errorStatus.Err()
}

func Errorf(c int, a ...interface{}) error {
	apiError, ok := APIErrorMap[c]
	if !ok {
		println("Wrong Code")
	}
	errorMessage := fmt.Sprintf(apiError.Message, a...)

	errResponseFormat := proto_gen.ErrorResponse{
		Code:       int64(apiError.Code),
		StatusCode: int64(apiError.StatusCode),
		Message:    errorMessage,
	}

	errorStatus, err := status.New(HttpCode2GrpcCode(apiError.Code), errorMessage).WithDetails(&errResponseFormat)

	if err != nil {
		println("Something Wrong")
	}

	return errorStatus.Err()
}
