package rm

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// error codes :
const (
	ErrorCodeTransactionNotFound              = "TRANSACTION_NOT_FOUND"
	ErrorCodeStoreNotFound                    = "STORE_NOT_FOUND"
	ErrorCodeMerchantSettlementAccNotVerified = "MERCHANT_SETTLEMENT_ACCOUNT_NOT_VERIFIED"
)

// Error :
type Error struct {
	Code        string
	Msg         string
	url         string
	rawRequest  []byte
	rawResponse []byte
}

var (
	_ fmt.Formatter = (*Error)(nil)
	_ error         = (*Error)(nil)
)

func newError(url string, reqBytes, respBytes []byte) *Error {
	e := new(Error)
	e.Code = strings.ToUpper(strings.TrimSpace(gjson.GetBytes(respBytes, "error.code").String()))
	e.Msg = strings.TrimSpace(gjson.GetBytes(respBytes, "error.message").String())
	e.url = url
	e.rawResponse = respBytes
	e.rawRequest = reqBytes
	return e
}

// Error :
func (e Error) Error() string {
	return "rm: " + string(e.Code)
}

func (e Error) Format(f fmt.State, verb rune) {
	if !f.Flag('+') {
		f.Write([]byte(e.Error()))
		return
	}

	f.Write([]byte("URL : "))
	f.Write([]byte(e.url))
	f.Write([]byte("\n"))
	f.Write([]byte("Request Body :\n"))
	f.Write(e.rawRequest)
	f.Write([]byte("\n"))
	f.Write([]byte("Response Body :\n"))
	f.Write(e.rawResponse)
}

func (e Error) Response() string {
	return string(e.rawResponse)
}

func (e Error) ResponseBytes() []byte {
	return e.rawResponse
}
