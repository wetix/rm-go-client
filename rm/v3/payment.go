package rm

import (
	"context"
	"errors"
)

type layout string

const (
	LayoutV1 layout = "v1"
	LayoutV2 layout = "v2"
	LayoutV3 layout = "v3"
)

type (
	PaymentType   string
	PaymentMethod string
)

const (
	PaymentTypeWeb    PaymentType = "WEB_PAYMENT"
	PaymentTypeMobile PaymentType = "MOBILE_PAYMENT"

	PaymentMethodGrab  PaymentMethod = "GRABPAY_MY"
	PaymentMethodTnG   PaymentMethod = "TNG_MY"
	PaymentMethodBoost PaymentMethod = "BOOST_MY"
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
	Type          PaymentType     `json:"type"`
	Method        []PaymentMethod `json:"method"`
	StoreID       string          `json:"storeId"`
	RedirectURL   string          `json:"redirectUrl"`
	NotifyURL     string          `json:"notifyUrl"`
	LayoutVersion layout          `json:"layoutVersion"`
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
	req.LayoutVersion = LayoutV2
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
		res, err := c.GetStores(ctx)
		if err != nil {
			return nil, err
		}

		if len(res.Items) == 0 {
			return nil, errors.New("no available store to init payment")
		}

		req.StoreID = res.Items[0].ID
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
