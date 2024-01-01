package rm

import (
	"context"
	"errors"
)

// CreatePaymentCheckoutRequest :
type CreatePaymentCheckoutRequest struct {
	Order struct {
		ID             string `json:"id"`
		Title          string `json:"title"`
		Detail         string `json:"detail"`
		AdditionalData string `json:"additionalData"`
		Amount         uint   `json:"amount"`
		Currency       string `json:"currencyType"`
	} `json:"order"`
	Customer struct {
		UserID      string `json:"userId"`
		Email       string `json:"email,omitempty"`
		CountryCode string `json:"countryCode,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
	} `json:"customer"`
	Type             PaymentType     `json:"type"`
	Method           []PaymentMethod `json:"method"`
	ExcludeMethod    []string        `json:"excludeMethod"`
	StoreID          string          `json:"storeId"`
	RedirectURL      string          `json:"redirectUrl"`
	NotifyURL        string          `json:"notifyUrl"`
	LayoutVersion    layout          `json:"layoutVersion"`
	ExpiresInSeconds int64           `json:"expiresInSeconds"`
}

// CreatePaymentCheckoutResponse :
type CreatePaymentCheckoutResponse struct {
	Item struct {
		CheckoutID string `json:"checkoutId"`
		URL        string `json:"url"`
	} `json:"item"`
	Code string `json:"code"`
}

// CreatePaymentCheckout :
func (c *Client) CreatePaymentCheckout(
	ctx context.Context,
	req CreatePaymentCheckoutRequest,
) (*CreatePaymentCheckoutResponse, error) {
	if req.Method == nil {
		req.Method = make([]PaymentMethod, 0)
	}
	if req.Type == "" {
		req.Type = PaymentTypeWeb
	}
	if req.Order.Currency == "" {
		req.Order.Currency = "MYR"
	}
	if req.StoreID == "" {
		if c.storeID != "" {
			req.StoreID = c.storeID
		} else {
			// prevent concurrency race
			c.mu.Lock()
			defer c.mu.Lock()
			res, err := c.GetStores(ctx)
			if err != nil {
				return nil, err
			}

			if len(res.Items) == 0 {
				return nil, errors.New("rm: no available store to init payment")
			}

			req.StoreID = res.Items[0].ID
		}
	}

	resp := new(CreatePaymentCheckoutResponse)
	if err := c.do(
		ctx,
		"create_payment_checkout",
		"post",
		c.openEndpoint+"/v3/payment/online",
		req,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
