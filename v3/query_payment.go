package rm

import (
	"context"
	"time"
)

// GetPaymentByOrderIDResponse :
type GetPaymentByOrderIDResponse struct {
	Item struct {
		Store struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			ImageURL     string `json:"imageUrl"`
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
		} `json:"store"`
		ReferenceID   string `json:"referenceId"`
		TransactionID string `json:"transactionId"`
		Order         struct {
			ID     string `json:"id"`
			Title  string `json:"title"`
			Detail string `json:"detail"`
			Amount int    `json:"amount"`
		} `json:"order"`
		TerminalID string `json:"terminalId"`
		Payee      struct {
		} `json:"payee"`
		CurrencyType  string        `json:"currencyType"`
		BalanceAmount int           `json:"balanceAmount"`
		Platform      string        `json:"platform"`
		Method        string        `json:"method"`
		TransactionAt time.Time     `json:"transactionAt"`
		Type          PaymentType   `json:"type"`
		Status        PaymentStatus `json:"status"`
		Region        string        `json:"region"`
		CreatedAt     time.Time     `json:"createdAt"`
		UpdatedAt     time.Time     `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
}

// GetPaymentByOrderID :
func (c *Client) GetPaymentByOrderID(
	ctx context.Context,
	orderID string,
) (*GetPaymentByOrderIDResponse, error) {
	resp := new(GetPaymentByOrderIDResponse)
	if err := c.do(
		ctx,
		"query_payment",
		"get",
		c.openEndpoint+"/v3/payment/transaction/order/"+orderID,
		nil,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
