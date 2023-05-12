package util

import (
	"bytes"
	"copilotiq/postgrid-client-golang/postgrid"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func TypeToReader(i interface{}) (*bytes.Reader, *postgrid.APIError) {
	jsonBytes, jsonErr := json.Marshal(i)
	if jsonErr != nil {
		return nil, postgrid.BuildError(500, fmt.Sprintf("error reading req [%+v]", i), "client_internal_error")
	}

	return bytes.NewReader(jsonBytes), nil
}

func ResToType(code int, reader io.Reader, successType interface{}) *postgrid.APIError {
	if code < http.StatusOK || (code < http.StatusBadRequest && code >= http.StatusMultipleChoices) {
		return postgrid.BuildError(500, fmt.Sprintf("unexpected status code [%d]", code), "client_implementation_error")
	}

	resBody, err := io.ReadAll(reader)
	if err != nil {
		return postgrid.BuildError(500, fmt.Sprintf("error reading response body [%+v] with err [%+v]", string(resBody), err), "client_receive_error")
	}

	var jsonErr error
	var isError bool
	var serverErr postgrid.APIError
	if code >= http.StatusBadRequest {
		isError = true
		jsonErr = json.Unmarshal(resBody, &serverErr)
	} else {
		jsonErr = json.Unmarshal(resBody, &successType)
	}

	if jsonErr != nil {
		return postgrid.BuildError(500, fmt.Sprintf("error unmarshalling res [%+v]", string(resBody)), "client_validation_error")
	}

	if isError {
		return &serverErr
	}

	return nil
}
