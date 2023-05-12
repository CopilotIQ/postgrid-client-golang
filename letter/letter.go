package letter

import (
	"copilotiq/postgrid-client-golang/postgrid"
	"copilotiq/postgrid-client-golang/util"
	"fmt"
	"net/http"
	"time"
)

const BaseUrl = "https://api.postgrid.com/print-mail/v1/letters"
const FirstClass = "first_class"
const StandardClass = "standard_class"

type MailingClass string
type MergeVariables map[string]interface{}

type CreateReq struct {
	Color          bool           `json:"color"`
	From           Contact        `json:"from"`
	MailingClass   MailingClass   `json:"mailingClass"`
	MergeVariables MergeVariables `json:"mergeVariables"`
	Template       string         `json:"template"`
	To             Contact        `json:"to"`
}

type CreateRes struct {
	AddressPlacement string         `json:"addressPlacement"`
	Color            bool           `json:"color"`
	CreatedAt        time.Time      `json:"createdAt"`
	DoubleSided      bool           `json:"doubleSided"`
	EnvelopeType     string         `json:"envelopeType"`
	From             Contact        `json:"from"`
	ID               string         `json:"id"`
	Live             bool           `json:"live"`
	MailingClass     string         `json:"mailingClass"`
	MergeVariables   MergeVariables `json:"mergeVariables"`
	Object           string         `json:"object"`
	SendDate         time.Time      `json:"sendDate"`
	Size             string         `json:"size"`
	Status           string         `json:"status"`
	Template         string         `json:"template"`
	To               Contact        `json:"to"`
	UpdatedAt        time.Time      `json:"updatedAt"`
}

func New(apiKey string) *Client {
	pg := postgrid.New(apiKey)
	return &Client{
		baseURL: BaseUrl,
		client:  http.DefaultClient,
		pg:      pg,
	}
}

type Client struct {
	baseURL string
	pg      *postgrid.PostGrid
	client  *http.Client
}

func (c *Client) Post(req *http.Request) (*http.Response, *postgrid.APIError) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("x-api-key", c.pg.APIKey())

	req.Header = headers
	res, err := c.client.Do(req)
	if err != nil {
		return nil, postgrid.BuildError(500, fmt.Sprintf("error sending req [%+v]", req), "client_transmit_error")
	}

	return res, nil
}

func (c *Client) CreateLetter(req *CreateReq) (*CreateRes, *postgrid.APIError) {
	bodyReader, typeErr := util.TypeToReader(req)
	if typeErr != nil {
		return nil, typeErr
	}

	postReq, err := http.NewRequest("POST", c.baseURL, bodyReader)
	if err != nil {
		return nil, postgrid.BuildError(500, fmt.Sprintf("error generating POST req [%+v] for req [%+v]", err, req), "client_internal_error")
	}

	res, postErr := c.Post(postReq)
	if postErr != nil {
		return nil, postErr
	}

	var createRes CreateRes
	resErr := util.ResToType(res.StatusCode, res.Body, &createRes)
	if resErr != nil {
		return nil, resErr
	}

	return &createRes, nil
}
