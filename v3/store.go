package rm

import (
	"context"
	"time"
)

// Store :
type Store struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PostCode     string `json:"postCode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	CountryCode  string `json:"countryCode"`
	PhoneNumber  string `json:"phoneNumber"`
	GeoLocation  struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geoLocation"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetStoresResponse :
type GetStoresResponse struct {
	Items []Store `json:"items"`
	Code  string  `json:"code"`
	Meta  struct {
		Count int `json:"count"`
		Total int `json:"total"`
	} `json:"meta"`
}

// GetStores :
func (c *Client) GetStores(ctx context.Context) (*GetStoresResponse, error) {
	resp := new(GetStoresResponse)
	if err := c.do(
		ctx,
		"get_stores",
		"get",
		c.openEndpoint+"/v3/stores?limit=100",
		nil,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
