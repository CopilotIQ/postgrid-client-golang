package letter

import (
	"encoding/json"
	"github.com/jgroeneveld/trial/assert"
	"github.com/joho/godotenv"
	"github.com/nsf/jsondiff"
	"log"
	"os"
	"reflect"

	"testing"
	"time"
)

const ApiKeyEnvKey = "POST_GRID_API_KEY"

var TestClient *Client

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Unable to load .env file: %s", err)
	}

	apiKey := os.Getenv(ApiKeyEnvKey)
	if apiKey == "" {
		log.Fatalf("Cannot proceed when apiKey is the empty string [%s]", apiKey)
	}

	TestClient = New(apiKey)
	if TestClient.pg.IsLive() {
		log.Fatalf("Cannot proceed when API key is live [%s]", apiKey)
	}
}

func TestCreateReq(t *testing.T) {
	expected :=
		`
{
    "from": {
        "firstName": "Four Seasons Hotel",
        "lastName": "LOS ANGELES At Beverly Hills",
        "addressLine1": "300 DOHENY DR",
        "addressLine2": "ROOM 1234",
        "city": "LOS ANGELES",
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
    "template": "template_eeiS9Gc4DyKyDqxSoKtdfw",
    "color": false,
    "mergeVariables": {
        "date": "May 1st, 2023",
        "greeting": "Hello GoLang,",
        "int": 42,
        "float": 92.12
    }
}
	`

	t.Run("verify CreateReq marshals to expected", func(t *testing.T) {
		jsonBytes, err := json.Marshal(GenerateCreateReq())
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
        "lastName": "LOS ANGELES At Beverly Hills",
        "postalOrZip": "90048",
        "provinceOrState": "CA"
    },
    "mailingClass": "standard_class",
    "mergeVariables": {
        "date": "May 1st, 2023",
        "float": 92.12,
        "greeting": "Hello GoLang,",
        "int": 42
    },
    "sendDate": "2023-05-11T18:52:36.68Z",
    "size": "us_letter",
    "status": "ready",
    "template": "template_eeiS9Gc4DyKyDqxSoKtdfw",
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

	t.Run("verify CreateRes marshals to expected", func(t *testing.T) {
		jsonBytes, err := json.Marshal(GenerateCreateRes())
		assert.Nil(t, err)
		opts := jsondiff.DefaultConsoleOptions()
		diff, _ := jsondiff.Compare([]byte(expected), jsonBytes, &opts)
		assert.Equal(t, int(jsondiff.FullMatch), int(diff))
	})
}

func TestCreate(t *testing.T) {
	t.Run("verify known response from known input", func(t *testing.T) {
		cReq := GenerateCreateReq()
		cRes, err := TestClient.CreateLetter(cReq)
		assert.Nil(t, err)
		VerifyCreateReqVsCreateRes(t, cReq, cRes)
	})
}

func GenerateCreateReq() *CreateReq {
	return &CreateReq{
		Color:        false,
		MailingClass: FirstClass,
		Template:     "template_eeiS9Gc4DyKyDqxSoKtdfw",
		From: Contact{
			AddressLine1:    "300 DOHENY DR",
			AddressLine2:    "ROOM 1234",
			City:            "LOS ANGELES",
			CountryCode:     "US",
			FirstName:       "Four Seasons Hotel",
			LastName:        "LOS ANGELES At Beverly Hills",
			PostalOrZip:     "90048",
			ProvinceOrState: "CA",
		},
		To: Contact{
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
			"int":      float64(42), // TODO(Canavan): response upgrades 42 to float64 which causes reflect.DeepEqual to break
		},
	}
}

func GenerateCreateRes() *CreateRes {
	return &CreateRes{
		AddressPlacement: "top_first_page",
		Color:            false,
		CreatedAt:        time.Date(2023, time.May, 11, 18, 52, 36, 684000000, time.UTC),
		DoubleSided:      false,
		EnvelopeType:     "standard_double_window",
		ID:               "letter_nUhyevBaQfMByda8bCmMSk",
		Live:             false,
		MailingClass:     StandardClass,
		Object:           "letter",
		SendDate:         time.Date(2023, time.May, 11, 18, 52, 36, 680000000, time.UTC),
		Size:             "us_letter",
		Status:           "ready",
		Template:         "template_eeiS9Gc4DyKyDqxSoKtdfw",
		UpdatedAt:        time.Date(2023, time.May, 11, 18, 52, 36, 684000000, time.UTC),
		MergeVariables: MergeVariables{
			"date":     "May 1st, 2023",
			"float":    92.12,
			"greeting": "Hello GoLang,",
			"int":      float64(42),
		},
		From: Contact{
			AddressLine1:    "300 DOHENY DR",
			AddressLine2:    "ROOM 1234",
			AddressStatus:   "verified",
			City:            "LOS ANGELES",
			Country:         "UNITED STATES",
			CountryCode:     "US",
			FirstName:       "Four Seasons Hotel",
			ID:              "contact_6abGybQegaSeQ5cknTy9yV",
			LastName:        "LOS ANGELES At Beverly Hills",
			Object:          "contact",
			PostalOrZip:     "90048",
			ProvinceOrState: "CA",
		},
		To: Contact{
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
}

func VerifyCreateReqVsCreateRes(t *testing.T, cReq *CreateReq, cRes *CreateRes) {
	assert.Equal(t, cReq.Color, cRes.Color)
	VerifyBeforeContactVsAfterContact(t, &cReq.From, &cRes.From)
	assert.Equal(t, string(cReq.MailingClass), cRes.MailingClass)
	assert.MustBeDeepEqual(t, cReq.MergeVariables, cRes.MergeVariables)
	assert.Equal(t, cReq.Template, cRes.Template)
	assert.True(t, reflect.DeepEqual(cReq.To, cRes.To))
}

func VerifyBeforeContactVsAfterContact(t *testing.T, beforeC *Contact, afterC *Contact) {
	assert.Equal(t, beforeC.AddressLine1, afterC.AddressLine1)
	assert.Equal(t, beforeC.AddressLine2, afterC.AddressLine2)
	assert.Equal(t, beforeC.City, afterC.City)
	assert.Equal(t, beforeC.CountryCode, afterC.CountryCode)
	assert.Equal(t, beforeC.FirstName, afterC.FirstName)
	assert.Equal(t, beforeC.LastName, afterC.LastName)
	assert.Equal(t, beforeC.PostalOrZip, afterC.PostalOrZip)
	assert.Equal(t, beforeC.ProvinceOrState, afterC.ProvinceOrState)

	assert.Equal(t, "", beforeC.AddressStatus)
	assert.Equal(t, "", beforeC.Country)
	assert.Equal(t, "", beforeC.ID)
	assert.Equal(t, "", beforeC.Object)
	assert.Equal(t, "UNITED STATES", afterC.Country)
	assert.Equal(t, "contact", afterC.Object)
	assert.Equal(t, 30, len(afterC.ID))
}
