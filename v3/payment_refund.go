package rm

import (
	"context"
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
