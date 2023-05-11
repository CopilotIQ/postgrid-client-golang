package letter

import (
	"copilotiq/postgrid-client-golang/contact"
	"encoding/json"
	"github.com/jgroeneveld/trial/assert"
	"github.com/nsf/jsondiff"

	"testing"
	"time"
)

func TestCreateReq(t *testing.T) {
	expected :=
		`
{
    "from": {
        "firstName": "Four Seasons Hotel",
        "lastName": "Los Angeles At Beverly Hills",
        "addressLine1": "300 Doheny Dr",
        "addressLine2": "Room 1234",
        "city": "Los Angeles",
        "countryCode": "US",
        "postalOrZip": "90048",
        "provinceOrState": "CA"
    },
    "to": {
        "firstName": "Mercedes-Benz",
        "lastName": "of Beverly Hills",
        "addressLine1": "9250 Beverly Blvd",
        "addressLine2": "Garage 32",
        "city": "Beverly Hills",
        "countryCode": "US",
        "postalOrZip": "90210",
        "provinceOrState": "CA"
    },
    "mailingClass": "first_class",
    "template": "template_12eCCX5GGG8cHfisGf6McD",
    "color": false,
    "mergeVariables": {
        "date": "May 1st, 2023",
        "greeting": "Hello GoLang,",
        "int": 42,
        "float": 92.12
    }
}
	`
	input := CreateReq{
		Color:        false,
		MailingClass: "first_class",
		Template:     "template_12eCCX5GGG8cHfisGf6McD",
		From: contact.Contact{
			AddressLine1:    "300 Doheny Dr",
			AddressLine2:    "Room 1234",
			City:            "Los Angeles",
			CountryCode:     "US",
			FirstName:       "Four Seasons Hotel",
			LastName:        "Los Angeles At Beverly Hills",
			PostalOrZip:     "90048",
			ProvinceOrState: "CA",
		},
		To: contact.Contact{
			AddressLine1:    "9250 Beverly Blvd",
			AddressLine2:    "Garage 32",
			City:            "Beverly Hills",
			CountryCode:     "US",
			FirstName:       "Mercedes-Benz",
			LastName:        "of Beverly Hills",
			PostalOrZip:     "90210",
			ProvinceOrState: "CA",
		},
		MergeVariables: MergeVariables{
			"date":     "May 1st, 2023",
			"float":    92.12,
			"greeting": "Hello GoLang,",
			"int":      42,
		},
	}

	t.Run("verify CreateReq marshals to expected", func(t *testing.T) {
		jsonBytes, err := json.Marshal(&input)
		assert.Nil(t, err)
		opts := jsondiff.DefaultConsoleOptions()
		diff, _ := jsondiff.Compare([]byte(expected), jsonBytes, &opts)
		assert.Equal(t, int(jsondiff.FullMatch), int(diff))
	})
}
func TestCreateRes(t *testing.T) {
	expected :=
		`
{
    "id": "letter_nUhyevBaQfMByda8bCmMSk",
    "object": "letter",
    "live": false,
    "addressPlacement": "top_first_page",
    "color": false,
    "doubleSided": false,
    "envelopeType": "standard_double_window",
    "from": {
        "id": "contact_6abGybQegaSeQ5cknTy9yV",
        "object": "contact",
        "addressLine1": "300 DOHENY DR",
        "addressLine2": "ROOM 1234",
        "addressStatus": "verified",
        "city": "LOS ANGELES",
        "country": "UNITED STATES",
        "countryCode": "US",
        "firstName": "Four Seasons Hotel",
        "lastName": "Los Angeles At Beverly Hills",
        "postalOrZip": "90048",
        "provinceOrState": "CA"
    },
    "mailingClass": "first_class",
    "mergeVariables": {
        "date": "May 1st, 2023",
        "float": 92.12,
        "greeting": "Hello GoLang,",
        "int": 42
    },
    "sendDate": "2023-05-11T18:52:36.68Z",
    "size": "us_letter",
    "status": "ready",
    "template": "template_12eCCX5GGG8cHfisGf6McD",
    "to": {
        "id": "contact_eKayBKrC356AZPNifvfrAL",
        "object": "contact",
        "addressLine1": "9250 BEVERLY BLVD",
        "addressLine2": "GARAGE 32",
        "addressStatus": "verified",
        "city": "BEVERLY HILLS",
        "country": "UNITED STATES",
        "countryCode": "US",
        "firstName": "Mercedes-Benz",
        "lastName": "of Beverly Hills",
        "postalOrZip": "90210",
        "provinceOrState": "CA"
    },
    "createdAt": "2023-05-11T18:52:36.684Z",
    "updatedAt": "2023-05-11T18:52:36.684Z"
}
`

	input := CreateRes{
		AddressPlacement: "top_first_page",
		Color:            false,
		CreatedAt:        time.Date(2023, time.May, 11, 18, 52, 36, 684000000, time.UTC),
		DoubleSided:      false,
		EnvelopeType:     "standard_double_window",
		ID:               "letter_nUhyevBaQfMByda8bCmMSk",
		Live:             false,
		MailingClass:     "first_class",
		Object:           "letter",
		SendDate:         time.Date(2023, time.May, 11, 18, 52, 36, 680000000, time.UTC),
		Size:             "us_letter",
		Status:           "ready",
		Template:         "template_12eCCX5GGG8cHfisGf6McD",
		UpdatedAt:        time.Date(2023, time.May, 11, 18, 52, 36, 684000000, time.UTC),
		MergeVariables: MergeVariables{
			"date":     "May 1st, 2023",
			"float":    92.12,
			"greeting": "Hello GoLang,",
			"int":      42,
		},
		From: contact.Contact{
			AddressLine1:    "300 DOHENY DR",
			AddressLine2:    "ROOM 1234",
			AddressStatus:   "verified",
			City:            "LOS ANGELES",
			Country:         "UNITED STATES",
			CountryCode:     "US",
			FirstName:       "Four Seasons Hotel",
			ID:              "contact_6abGybQegaSeQ5cknTy9yV",
			LastName:        "Los Angeles At Beverly Hills",
			Object:          "contact",
			PostalOrZip:     "90048",
			ProvinceOrState: "CA",
		},
		To: contact.Contact{
			AddressLine1:    "9250 BEVERLY BLVD",
			AddressLine2:    "GARAGE 32",
			AddressStatus:   "verified",
			City:            "BEVERLY HILLS",
			Country:         "UNITED STATES",
			CountryCode:     "US",
			FirstName:       "Mercedes-Benz",
			ID:              "contact_eKayBKrC356AZPNifvfrAL",
			LastName:        "of Beverly Hills",
			Object:          "contact",
			PostalOrZip:     "90210",
			ProvinceOrState: "CA",
		},
	}
	t.Run("verify CreateRes marshals to expected", func(t *testing.T) {
		jsonBytes, err := json.Marshal(&input)
		assert.Nil(t, err)
		opts := jsondiff.DefaultConsoleOptions()
		diff, _ := jsondiff.Compare([]byte(expected), jsonBytes, &opts)
		assert.Equal(t, int(jsondiff.FullMatch), int(diff))
	})
}
