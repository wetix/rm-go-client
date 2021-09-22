package rm

import (
	"context"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/require"
)

func mockRmClient() *Client {
	pk, _ := ioutil.ReadFile("../test/pk.pem")
	pub, _ := ioutil.ReadFile("../test/server_pub.pem")

	storeID := "2808912573238362402"
	return NewClient(
		Config{
			ClientID:     "1599646279297591629",
			ClientSecret: "NekiDbnNHbHLWdRmbqtwBCqywfYkVVnE",
			PrivateKey:   pk,
			StoreID:      storeID,
			PublicKey:    pub,
			Sandbox:      true,
		},
	)
}

func TestRmClient(t *testing.T) {
	ctx := context.Background()
	client := mockRmClient()

	req := CreatePaymentCheckoutRequest{}
	req.Order.ID = uniuri.NewLen(10)
	req.Order.Title = "Testing #" + req.Order.ID
	req.Order.Amount = 1000
	req.Customer.UserID = "1234"
	req.StoreID = client.storeID
	req.NotifyURL = "https://www.google.com"
	req.RedirectURL = "https://www.google.com"
	resp, err := client.GetStores(ctx)
	require.NoError(t, err)
	require.True(t, len(resp.Items) > 0)

	res, err := client.CreatePaymentCheckout(ctx, req)
	require.NoError(t, err)
	require.NotEmpty(t, res.Item.CheckoutID)
	require.NotEmpty(t, res.Item.URL)

	order, err := client.GetPaymentByCheckoutID(ctx, res.Item.CheckoutID)
	require.NoError(t, err)
	require.Equal(t, ResponseSuccess, order.Code)
	require.Equal(t, req.Order.ID, order.Item.Order.ID)
	require.Equal(t, res.Item.CheckoutID, order.Item.ID)
	require.Equal(t, req.Order.Title, order.Item.Order.Title)
	require.Equal(t, req.Order.Amount, order.Item.Order.Amount)
	require.Equal(t, req.NotifyURL, order.Item.NotifyURL)
	require.Equal(t, req.RedirectURL, order.Item.RedirectURL)

	{
		pymt, err := client.GetPaymentByOrderID(ctx, "128200910090623482313")
		require.NoError(t, err)
		require.Equal(t, ResponseSuccess, pymt.Code)
	}

	// invalid order id
	{
		_, err := client.GetPaymentByOrderID(ctx, "xxx")
		require.Error(t, err)
		require.True(t, errors.Is(err, ErrTransactionNotFound))
	}

	{
		req := RefundPaymentRequest{}
		req.TransactionID = "210922094709010322801968"
		req.Reason = "Not received goods"
		_, err := client.RefundPayment(ctx, req)
		require.True(t, errors.Is(err, ErrPaymentAlreadyRefunded))
	}

	{
		req := CreateTransactionQRRequest{}
		req.Amount = 10000
		req.Type = "STATIC"
		req.Order.Title = "Payment Test"
		req.CurrencyType = "MYR"
		req.IsPreFillAmount = false
		req.Method = make([]string, 0)
		req.RedirectURL = "www.google.com"
		req.Expiry.Type = "PERMANENT"
		req.StoreID = client.storeID
		resp, err := client.CreateTransactionQR(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, resp.Item.QrCodeURL)
	}

}
