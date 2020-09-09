### RevenueMonster Golang Client

Apart from the official [rm-sdk-go](https://github.com/RevenueMonster/rm-sdk-go) doesn't support `context.Context` and `jaeger` tracing, and this repository is mainly to cover this two core elements.


```go
import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/dchest/uniuri"
	rm "github.com/si3nloong/rm-go-client/rm/v3"
)

func main() {
    ctx := context.Background()
    pk, _ := ioutil.ReadFile("../test/pk.pem")
    pub, _ := ioutil.ReadFile("../test/server_pub.pem")

    client := rm.NewClient(
        "1599646279297591629",
        "NekiDbnNHbHLWdRmbqtwBCqywfYkVVnE",
        pk,
        pub,
        true,
    )

    req := rm.CreatePaymentCheckoutRequest{}
    req.Order.ID = uniuri.NewLen(10)
    req.Order.Title = "Testing #" + req.Order.ID
    req.Order.Amount = 1000
    req.NotifyURL = "https://www.google.com"
    req.RedirectURL = "https://www.google.com"

    res, err := client.CreatePaymentCheckout(ctx, req)
    if err != nil {
        panic(err)
    }
}
```