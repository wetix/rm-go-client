package rm

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebhook(t *testing.T) {
	var (
		ctx    = context.Background()
		client = mockRmClient()
	)

	f, err := os.OpenFile("./sample/webhook.json", os.O_RDWR, 0644)
	require.NoError(t, err)

	wh, err := client.VerifyWebhook(ctx, f)
	require.NoError(t, err)

	require.Equal(t, 2750, wh.Data.BalanceAmount)
	require.Equal(t, "128200910090623482313", wh.Data.Order.ID)
	require.Equal(t, "MYR", wh.Data.CurrencyType)
	require.Equal(t, "2009106165088944", wh.Data.ReferenceID)
	require.Equal(t, ResponseSuccess, wh.Data.Status)
	require.Equal(t, "1597245150176673383", wh.Data.Store.ID)
	require.Equal(t, PaymentTypeWeb, wh.Data.Type)

	// file content shouldn't empty
	_, exp := ioutil.ReadAll(f)
	require.NoError(t, exp)

	_, err = client.VerifyWebhook(ctx, f)
	require.Equal(t, io.EOF, err)

	buf := bytes.NewBufferString(`{"data": "error"}`)
	_, err = client.VerifyWebhook(ctx, buf)
	require.Error(t, err)
}
