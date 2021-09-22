<p align="center">
    <a href="https://github.com/si3nloong/rm-go-client/releases"><img src="https://img.shields.io/github/v/tag/si3nloong/rm-go-client" alt="semver tag" title="semver tag"/></a>
    <a href="https://goreportcard.com/report/github.com/si3nloong/rm-go-client"><img src="https://goreportcard.com/badge/github.com/si3nloong/rm-go-client" alt="go report card" title="go report card"/></a>
    <a href="https://github.com/si3nloong/rm-go-client/blob/main/LICENSE"><img src="https://img.shields.io/github/license/si3nloong/rm-go-client" alt="license" title="license"/></a>
</p>

# RevenueMonster Go Client

> Simplified version of rm sdk.

Apart from the official [rm-sdk-go](https://github.com/RevenueMonster/rm-sdk-go) doesn't support `context.Context` and [opentracing](https://github.com/opentracing/opentracing-go), and this repository is mainly to cover this two core elements.

## ✨ Features

- simple and mininal (no extra functions other than payment)
- support [Jaeger](https://www.jaegertracing.io/)
- support `context.Context`
- support decentralize access token store

## 📝 How to use?

```go
import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/dchest/uniuri"
	rm "github.com/si3nloong/rm-go-client/v3"
)

func main() {
    ctx := context.Background()
    pk, _ := ioutil.ReadFile("../test/pk.pem")
    pub, _ := ioutil.ReadFile("../test/server_pub.pem")

    client := rm.NewClient(
		rm.Config{
			ClientID:     "1599646279297591629",
			ClientSecret: "NekiDbnNHbHLWdRmbqtwBCqywfYkVVnE",
			PrivateKey:   pk,
			PublicKey:    pub,
			Sandbox:      true, // determine whether it's using sandbox environment
		},
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

## 📄 License

[MIT](https://github.com/si3nloong/rm-go-client/blob/main/LICENSE)

Copyright (c) 2021-present, SianLoong Lee