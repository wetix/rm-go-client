package rm

import (
	"context"
	"time"
)

// CreateTransactionQRType :
type CreateTransactionQRType string

const (
	CreateTransactionQRTypeStatic CreateTransactionQRType = "STATIC"
)

// CreateTransactionQRRequest :
type CreateTransactionQRRequest struct {
	Type            CreateTransactionQRType `json:"type"`
	CurrencyType    string                  `json:"currencyType"`
	Amount          int                     `json:"amount"`
	IsPreFillAmount bool                    `json:"isPreFillAmount"`
	Method          []string                `json:"method"`
	Order           struct {
		AdditionalData string `json:"additionalData"`
		Details        string `json:"details"`
		Title          string `json:"title"`
	} `json:"order"`
	Expiry struct {
		Type string `json:"type"`
	} `json:"expiry"`
	RedirectURL string `json:"redirectUrl"`
	StoreID     string `json:"storeId"`
}

// CreateTransactionQRResponse :
type CreateTransactionQRResponse struct {
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
		Type            string      `json:"type"`
		IsPreFillAmount bool        `json:"isPreFillAmount"`
		CurrencyType    string      `json:"currencyType"`
		Amount          int         `json:"amount"`
		Platform        string      `json:"platform"`
		Method          interface{} `json:"method"`
		Expiry          struct {
			Type      string    `json:"type"`
			Day       int       `json:"day"`
			ExpiredAt time.Time `json:"expiredAt"`
		} `json:"expiry"`
		Code        string `json:"code"`
		Status      string `json:"status"`
		QrCodeURL   string `json:"qrCodeUrl"`
		RedirectURL string `json:"redirectUrl"`
		Order       struct {
			Title          string `json:"title"`
			Detail         string `json:"detail"`
			AdditionalData string `json:"additionalData"`
		} `json:"order"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
}

// CreateTransactionQR :
func (c *Client) CreateTransactionQR(
	ctx context.Context,
	req CreateTransactionQRRequest,
) (*CreateTransactionQRResponse, error) {
	if req.CurrencyType == "" {
		req.CurrencyType = "MYR"
	}
	resp := new(CreateTransactionQRResponse)
	if err := c.do(
		ctx,
		"create_transaction_qrcode",
		"post",
		c.openEndpoint+"/v3/payment/transaction/qrcode",
		req,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
