package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APIError struct {
	Code   int           `json:"code"`
	Error  *ErrorDetails `json:"error"`
	Object string        `json:"object"`
}

type ErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func BuildError(code int, errorMessage, errorType string) *APIError {
	return &APIError{
		Code: code,
		Error: &ErrorDetails{
			Message: errorMessage,
			Type:    errorType,
		},
		Object: "error",
	}
}

func TypeToReader(i interface{}) (*bytes.Reader, *APIError) {
	jsonBytes, jsonErr := json.Marshal(i)
	if jsonErr != nil {
		return nil, &APIError{
			Code: 500,
			Error: &ErrorDetails{
				Message: fmt.Sprintf("error reading req [%+v]", i),
				Type:    "client_internal_error",
			},
			Object: "",
		}
	}

	return bytes.NewReader(jsonBytes), nil
}

func ResToType(code int, reader io.Reader, successType interface{}) *APIError {
	if code < http.StatusOK || (code < http.StatusBadRequest && code >= http.StatusMultipleChoices) {
		return &APIError{
			Code: 500,
			Error: &ErrorDetails{
				Message: fmt.Sprintf("unexpected status code [%d]", code),
				Type:    "client_implementation_error",
			},
			Object: "error",
		}
	}

	resBody, err := io.ReadAll(reader)
	if err != nil {
		return &APIError{
			Code: 500,
			Error: &ErrorDetails{
				Message: fmt.Sprintf("error reading response body [%+v] with err [%+v]", string(resBody), err),
				Type:    "client_receive_error",
			},
			Object: "error",
		}
	}

	var jsonErr error
	var isError bool
	var serverErr APIError
	if code >= http.StatusBadRequest {
		isError = true
		jsonErr = json.Unmarshal(resBody, &serverErr)
	} else {
		jsonErr = json.Unmarshal(resBody, &successType)
	}

	if jsonErr != nil {
		return &APIError{
			Code: 500,
			Error: &ErrorDetails{
				Message: fmt.Sprintf("error unmarshalling res [%+v]", string(resBody)),
				Type:    "client_validation_error",
			},
			Object: "error",
		}
	}

	if isError {
		return &serverErr
	}

	return nil
}
