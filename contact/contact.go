package contact

type Contact struct {
	ID              string `json:"id,omitempty"`
	Object          string `json:"object,omitempty"`
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	AddressStatus   string `json:"addressStatus,omitempty"`
	City            string `json:"city"`
	Country         string `json:"country,omitempty"`
	CountryCode     string `json:"countryCode"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName,omitempty"`
	PostalOrZip     string `json:"postalOrZip"`
	ProvinceOrState string `json:"provinceOrState"`
}
