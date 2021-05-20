package rm

import (
	"context"
	"errors"
)

type layout string

// layout versions :
const (
	LayoutV1 layout = "v1"
	LayoutV2 layout = "v2"
	LayoutV3 layout = "v3"
)

// types :
type (
	PaymentType   string
	PaymentMethod string
)

// payment types :
const (
	PaymentTypeWeb    PaymentType = "WEB_PAYMENT"
	PaymentTypeMobile PaymentType = "MOBILE_PAYMENT"

	PaymentMethodWeChatMalaysia    PaymentMethod = "WECHAT_MY"
	PaymentMethodWeChatChina       PaymentMethod = "WECHAT_CN"
	PaymentMethodBoostMalaysia     PaymentMethod = "BOOST_MY"
	PaymentMethodPrestoMalaysia    PaymentMethod = "PRESTO_MY"
	PaymentMethodAlipayChina       PaymentMethod = "ALIPAY_CN"
	PaymentMethodTnGMalaysia       PaymentMethod = "TNG_MY"
	PaymentMethodGrabMalaysia      PaymentMethod = "GRABPAY_MY"
	PaymentMethodMaybankMalaysia   PaymentMethod = "MAYBANK_MY"
	PaymentMethodRazerPayMalaysia  PaymentMethod = "RAZERPAY_MY"
	PaymentMethodMCashMalaysia     PaymentMethod = "MCASH_MY"
	PaymentMethodShopeePayMalaysia PaymentMethod = "SHOPEEPAY_MY"
	PaymentMethodFpxMalaysia       PaymentMethod = "FPX_MY"
	PaymentMethodGoBizMalaysia     PaymentMethod = "GOBIZ_MY"
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
	req.LayoutVersion = LayoutV3
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
		if c.storeID == "" {
			res, err := c.GetStores(ctx)
			if err != nil {
				return nil, err
			}

			if len(res.Items) == 0 {
				return nil, errors.New("no available store to init payment")
			}

			req.StoreID = res.Items[0].ID
		} else {
			req.StoreID = c.storeID
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
