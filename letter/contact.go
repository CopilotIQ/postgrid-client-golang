package letter

type Contact struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	AddressStatus   string `json:"addressStatus,omitempty"`
	City            string `json:"city"`
	Country         string `json:"country,omitempty"`
	CountryCode     string `json:"countryCode"`
	FirstName       string `json:"firstName"`
	ID              string `json:"id,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	Object          string `json:"object,omitempty"`
	PostalOrZip     string `json:"postalOrZip"`
	ProvinceOrState string `json:"provinceOrState"`
}
