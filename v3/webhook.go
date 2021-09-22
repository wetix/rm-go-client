package rm

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"
)

type eventType string

const (
	EventTypeWebPayment eventType = "PAYMENT_WEB_ONLINE"
)

type Webhook struct {
	Data struct {
		BalanceAmount int       `json:"balanceAmount"`
		CreatedAt     time.Time `json:"createdAt"`
		CurrencyType  string    `json:"currencyType"`
		Method        string    `json:"method"`
		Order         struct {
			Amount int    `json:"amount"`
			Detail string `json:"detail"`
			ID     string `json:"id"`
			Title  string `json:"title"`
		} `json:"order"`
		Payee struct {
		} `json:"payee"`
		Platform    string `json:"platform"`
		ReferenceID string `json:"referenceId"`
		Region      string `json:"region"`
		Status      string `json:"status"`
		Store       struct {
			AddressLine1 string    `json:"addressLine1"`
			AddressLine2 string    `json:"addressLine2"`
			City         string    `json:"city"`
			Country      string    `json:"country"`
			CountryCode  string    `json:"countryCode"`
			CreatedAt    time.Time `json:"createdAt"`
			GeoLocation  struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"geoLocation"`
			ID          string    `json:"id"`
			ImageURL    string    `json:"imageUrl"`
			Name        string    `json:"name"`
			PhoneNumber string    `json:"phoneNumber"`
			PostCode    string    `json:"postCode"`
			State       string    `json:"state"`
			Status      string    `json:"status"`
			UpdatedAt   time.Time `json:"updatedAt"`
		} `json:"store"`
		TerminalID    string      `json:"terminalId"`
		TransactionAt time.Time   `json:"transactionAt"`
		TransactionID string      `json:"transactionId"`
		Type          PaymentType `json:"type"`
		UpdatedAt     time.Time   `json:"updatedAt"`
		Voucher       interface{} `json:"voucher"`
	} `json:"data"`
	EventType eventType `json:"eventType"`
}

// VerifyWebhook :
func (c *Client) VerifyWebhook(
	ctx context.Context,
	body io.Reader,
) (*Webhook, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return nil, err
	}

	wh := new(Webhook)
	if err := json.NewDecoder(buf).Decode(wh); err != nil {
		return nil, err
	}
	return wh, nil
}
