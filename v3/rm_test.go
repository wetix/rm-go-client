package rm

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/require"
)

func TestRMClient(t *testing.T) {
	ctx := context.Background()
	pk, _ := ioutil.ReadFile("../test/pk.pem")
	pub, _ := ioutil.ReadFile("../test/server_pub.pem")

	storeID := "2808912573238362402"
	client := NewClient(
		Config{
			ClientID:     "1599646279297591629",
			ClientSecret: "NekiDbnNHbHLWdRmbqtwBCqywfYkVVnE",
			PrivateKey:   pk,
			PublicKey:    pub,
			Sandbox:      true,
		},
	)

	req := CreatePaymentCheckoutRequest{}
	req.Order.ID = uniuri.NewLen(10)
	req.Order.Title = "Testing #" + req.Order.ID
	req.Order.Amount = 1000
	req.Customer.UserID = "1234"
	req.StoreID = storeID
	req.NotifyURL = "https://www.google.com"
	req.RedirectURL = "https://www.google.com"
	resp, err := client.GetStores(ctx)
	require.NoError(t, err)
	require.True(t, len(resp.Items) > 0)

	res, err := client.CreatePaymentCheckout(ctx, req)
	require.NoError(t, err)
	require.NotEmpty(t, res.Item.CheckoutID)
	require.NotEmpty(t, res.Item.URL)

	pymt, err := client.GetPaymentByOrderID(ctx, "128200910090623482313")
	require.NoError(t, err)
	require.Equal(t, "SUCCESS", pymt.Code)

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
		req.StoreID = storeID
		resp, err := client.CreateTransactionQR(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, resp.Item.QrCodeURL)
	}
}
