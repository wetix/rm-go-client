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
			Amount uint   `json:"amount"`
		} `json:"order"`
		TerminalID string `json:"terminalId"`
		Payee      struct {
		} `json:"payee"`
		CurrencyType  string        `json:"currencyType"`
		BalanceAmount uint          `json:"balanceAmount"`
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
		"query_payment_by_order_id",
		"get",
		c.openEndpoint+"/v3/payment/transaction/order/"+orderID,
		nil,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetPaymentByTransactionIDResponse :
type GetPaymentByTransactionIDResponse struct {
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
			Amount uint   `json:"amount"`
		} `json:"order"`
		TerminalID string `json:"terminalId"`
		Payee      struct {
		} `json:"payee"`
		CurrencyType  string        `json:"currencyType"`
		BalanceAmount int           `json:"balanceAmount"`
		Voucher       interface{}   `json:"voucher"`
		Platform      string        `json:"platform"`
		Method        string        `json:"method"`
		TransactionAt time.Time     `json:"transactionAt"`
		Type          string        `json:"type"`
		Status        PaymentStatus `json:"status"`
		Region        string        `json:"region"`
		Source        string        `json:"source"`
		CreatedAt     time.Time     `json:"createdAt"`
		UpdatedAt     time.Time     `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
}

// GetPaymentByTransactionID :
func (c *Client) GetPaymentByTransactionID(
	ctx context.Context,
	transactionID string,
) (*GetPaymentByTransactionIDResponse, error) {
	resp := new(GetPaymentByTransactionIDResponse)
	if err := c.do(
		ctx,
		"query_payment_by_transaction_id",
		"get",
		c.openEndpoint+"/v3/payment/transaction/"+transactionID,
		nil,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetPaymentByCheckoutIDResponse struct {
	Item struct {
		ID    string `json:"id"`
		Order struct {
			ID             string `json:"id"`
			Title          string `json:"title"`
			Detail         string `json:"detail"`
			AdditionalData string `json:"additionalData"`
			CurrencyType   string `json:"currencyType"`
			Amount         uint   `json:"amount"`
		} `json:"order"`
		Type          string    `json:"type"`
		TransactionID string    `json:"transactionId"`
		Platform      string    `json:"platform"`
		Method        []string  `json:"method"`
		RedirectURL   string    `json:"redirectUrl"`
		NotifyURL     string    `json:"notifyUrl"`
		StartAt       time.Time `json:"startAt"`
		EndAt         time.Time `json:"endAt"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
}

// GetPaymentByCheckoutID :
func (c *Client) GetPaymentByCheckoutID(
	ctx context.Context,
	checkoutID string,
) (*GetPaymentByCheckoutIDResponse, error) {
	resp := new(GetPaymentByCheckoutIDResponse)
	if err := c.do(
		ctx,
		"query_payment_by_checkout_id",
		"get",
		c.openEndpoint+"/v3/payment/online?checkoutId="+checkoutID,
		nil,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
