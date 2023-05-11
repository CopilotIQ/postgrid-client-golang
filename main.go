package main

import (
	"copilotiq/postgrid-client-golang/letter"
	"strings"
)

type PostGrid struct {
	apiKey string
	live   bool
}

type APIError struct {
	Code   int          `json:"code"`
	Error  ErrorDetails `json:"error"`
	Object string       `json:"object"`
}

type ErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
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

func (pg *PostGrid) IsLive() bool {
	return pg.live
}

func (pg *PostGrid) CreateLetter(req letter.CreateReq) (*letter.CreateRes, *APIError) {
	return nil, nil
}
