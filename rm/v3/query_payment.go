package rm

import (
	"context"
	"log"
)

// GetPaymentByOrderIDResponse :
type GetPaymentByOrderIDResponse struct {
}

// GetPaymentByOrderID :
func (c *Client) GetPaymentByOrderID(
	ctx context.Context,
	orderID string,
) (*GetPaymentByOrderIDResponse, error) {
	resp := new(GetPaymentByOrderIDResponse)
	log.Println(c.openEndpoint + "/v3/payment/transaction/order/" + orderID)
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
