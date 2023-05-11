package letter

import (
	"copilotiq/postgrid-client-golang/contact"
	"time"
)

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
