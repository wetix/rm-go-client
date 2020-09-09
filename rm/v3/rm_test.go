package rm

import (
	"context"
	"io/ioutil"
	"log"
	"testing"

	"github.com/dchest/uniuri"
	"github.com/stretchr/testify/require"
)

func TestRMClient(t *testing.T) {
	ctx := context.Background()
	pk, _ := ioutil.ReadFile("../test/pk.pem")
	pub, _ := ioutil.ReadFile("../test/server_pub.pem")

	client := NewClient(
		"1599646279297591629",
		"NekiDbnNHbHLWdRmbqtwBCqywfYkVVnE",
		pk,
		pub,
		true,
	)

	req := CreatePaymentCheckoutRequest{}
	req.Order.ID = uniuri.NewLen(10)
	req.Order.Title = "Testing #" + req.Order.ID
	req.Order.Amount = 1000
	req.NotifyURL = "https://www.google.com"
	req.RedirectURL = "https://www.google.com"
	resp, err := client.GetStores(ctx)
	require.NoError(t, err)
	require.True(t, len(resp.Items) > 0)

	res, err := client.CreatePaymentCheckout(ctx, req)
	require.NoError(t, err)
	require.NotEmpty(t, res.Item.CheckoutID)
	require.NotEmpty(t, res.Item.URL)

	pymt, err := client.GetPaymentByOrderID(ctx, req.Order.ID)
	require.NoError(t, err)
	log.Println(pymt)
}
