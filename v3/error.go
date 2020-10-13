package rm

import (
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
	Code string
	Msg  string
	raw  []byte
}

func newError(b []byte) *Error {
	e := new(Error)
	e.Code = strings.ToUpper(strings.TrimSpace(gjson.GetBytes(b, "error.code").String()))
	e.Msg = strings.TrimSpace(gjson.GetBytes(b, "error.message").String())
	e.raw = b
	return e
}

// Error :
func (e Error) Error() string {
	return "rm: " + string(e.Code)
}

func (e Error) Raw() string {
	return string(e.raw)
}
