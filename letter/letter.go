package letter

import (
	"copilotiq/postgrid-client-golang/contact"
	"copilotiq/postgrid-client-golang/impl"
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
	Color          bool            `json:"color"`
	From           contact.Contact `json:"from"`
	MailingClass   MailingClass    `json:"mailingClass"`
	MergeVariables MergeVariables  `json:"mergeVariables"`
	Template       string          `json:"template"`
	To             contact.Contact `json:"to"`
}

type CreateRes struct {
	AddressPlacement string          `json:"addressPlacement"`
	Color            bool            `json:"color"`
	CreatedAt        time.Time       `json:"createdAt"`
	DoubleSided      bool            `json:"doubleSided"`
	EnvelopeType     string          `json:"envelopeType"`
	From             contact.Contact `json:"from"`
	ID               string          `json:"id"`
	Live             bool            `json:"live"`
	MailingClass     string          `json:"mailingClass"`
	MergeVariables   MergeVariables  `json:"mergeVariables"`
	Object           string          `json:"object"`
	SendDate         time.Time       `json:"sendDate"`
	Size             string          `json:"size"`
	Status           string          `json:"status"`
	Template         string          `json:"template"`
	To               contact.Contact `json:"to"`
	UpdatedAt        time.Time       `json:"updatedAt"`
}

func New(apiKey string) *Client {
	pg := impl.New(apiKey)
	return &Client{
		baseURL: BaseUrl,
		pg:      pg,
	}
}

type Client struct {
	baseURL string
	pg      *impl.PostGrid
}

func (c *Client) CreateLetter(req *CreateReq) (*CreateRes, *util.APIError) {
	bodyReader, typeErr := util.TypeToReader(req)
	if typeErr != nil {
		return nil, typeErr
	}

	postReq, err := http.NewRequest("POST", c.baseURL, bodyReader)
	if err != nil {
		return nil, util.BuildError(500, fmt.Sprintf("error generating POST req [%+v] for req [%+v]", err, req), "client_internal_error")
	}

	ctHeader := http.Header{}
	postReq.Header["Content-Type"] = "application/json"

	resp, postErr := http.Post(c.baseURL, "application/json", bodyReader)
	if postErr != nil {
		return nil, util.BuildError(500, fmt.Sprintf("error sending req [%+v]", req), "client_transmit_error")
	}

	var createRes CreateRes
	resErr := util.ResToType(resp.StatusCode, resp.Body, &createRes)
	if resErr != nil {
		return nil, resErr
	}

	return &createRes, nil
}
