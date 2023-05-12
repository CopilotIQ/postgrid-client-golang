package postgrid

import "strings"

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

func New(apiKey string) *PostGrid {
	live := false

	if strings.HasPrefix(apiKey, "live_") {
		live = true
	}

	return &PostGrid{
		apiKey: apiKey,
		live:   live,
	}
}

type PostGrid struct {
	apiKey  string
	baseURL string
	live    bool
}

func (pg *PostGrid) APIKey() string {
	return pg.apiKey
}

func (pg *PostGrid) IsLive() bool {
	return pg.live
}
