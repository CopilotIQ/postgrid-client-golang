package contact

import (
	"encoding/json"
	"github.com/jgroeneveld/trial/assert"
	"github.com/nsf/jsondiff"
	"testing"
)

func TestContact(t *testing.T) {
	t.Run("verify to contact", func(t *testing.T) {
		expected := `
    {
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
    }
`

		from := Contact{
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
		}

		jsonBytes, err := json.Marshal(&from)
		assert.Nil(t, err)
		opts := jsondiff.DefaultConsoleOptions()
		diff, _ := jsondiff.Compare([]byte(expected), jsonBytes, &opts)
		assert.Equal(t, int(jsondiff.FullMatch), int(diff))
	})
	t.Run("verify from contact", func(t *testing.T) {
		expected := `
    {
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
    }
`

		to := Contact{
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
		}

		jsonBytes, err := json.Marshal(&to)
		assert.Nil(t, err)
		opts := jsondiff.DefaultConsoleOptions()
		diff, _ := jsondiff.Compare([]byte(expected), jsonBytes, &opts)
		assert.Equal(t, int(jsondiff.FullMatch), int(diff))
	})
}
