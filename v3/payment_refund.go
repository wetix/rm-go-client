package rm

import "context"

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
}

// RefundPayment :
func (c *Client) RefundPayment(
	ctx context.Context,
	orderID string,
	reason string,
) (*RefundPaymentResponse, error) {
	req := RefundPaymentRequest{}
	// req.Type = "FULL"
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
