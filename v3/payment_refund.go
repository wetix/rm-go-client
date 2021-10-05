package rm

import (
	"context"
	"time"
)

// RefundPaymentRequest :
type RefundPaymentRequest struct {
	TransactionID string `json:"transactionId"`
	Refund        struct {
		Type         string `json:"type"`
		CurrencyType string `json:"currencyType"`
		Amount       uint   `json:"amount"`
	} `json:"refund"`
	Reason string `json:"reason"`
}

// RefundPaymentResponse :
type RefundPaymentResponse struct {
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
			UserID string `json:"userId"`
		} `json:"payee"`
		CurrencyType  string      `json:"currencyType"`
		BalanceAmount int         `json:"balanceAmount"`
		Voucher       interface{} `json:"voucher"`
		Platform      string      `json:"platform"`
		Method        string      `json:"method"`
		TransactionAt time.Time   `json:"transactionAt"`
		Type          string      `json:"type"`
		Status        string      `json:"status"`
		Region        string      `json:"region"`
		ExtraInfo     struct {
			Card struct {
			} `json:"card"`
		} `json:"extraInfo"`
		Source    string    `json:"source"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"item"`
	Code string `json:"code"`
}

// RefundPayment :
func (c *Client) RefundPayment(
	ctx context.Context,
	req RefundPaymentRequest,
) (*RefundPaymentResponse, error) {
	pymt, err := c.GetPaymentByTransactionID(ctx, req.TransactionID)
	if err != nil {
		return nil, err
	}

	// if pymt.Item.Status != PaymentStatusSuccess {
	// 	return nil,
	// }

	// if amount is zero, we will perform full refunded
	if req.Refund.Amount == 0 {
		req.Refund.Type = "FULL"
		// FIXME: support partial refund
		// if orderResp.Item.Order.Amount < orderResp.Item.BalanceAmount {
		// }
		req.Refund.Amount = pymt.Item.Order.Amount
	}

	if req.Refund.CurrencyType == "" {
		req.Refund.CurrencyType = "MYR"
	}

	resp := new(RefundPaymentResponse)
	if err := c.do(
		ctx,
		"refund_payment",
		"post",
		c.openEndpoint+"/v3/payment/refund",
		req,
		resp,
	); err != nil {
		return nil, err
	}
	return resp, nil
}
