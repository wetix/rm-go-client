package rm

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// error codes :
const (
	errTemplate = "rm: %s"

	ErrorCodePaymentAlreadyRefunded           = "PAYMENT_FULLY_REFUNDED"
	ErrorCodeTransactionNotFound              = "TRANSACTION_NOT_FOUND"
	ErrorCodeStoreNotFound                    = "STORE_NOT_FOUND"
	ErrorCodeValidationError                  = "VALIDATION_ERROR"
	ErrorCodeRefundAmountExceedPerDay         = "PAYMENT_REFUND_AMOUNT_EXCEED_PER_DAY"
	ErrorCodeMerchantSettlementAccNotVerified = "MERCHANT_SETTLEMENT_ACCOUNT_NOT_VERIFIED"
)

// error to compare, you may compare using errors.Is(err, rm.ErrPaymentAlreadyRefunded)
var (
	ErrPaymentAlreadyRefunded  = newErrorCode(ErrorCodePaymentAlreadyRefunded)
	ErrTransactionNotFound     = newErrorCode(ErrorCodeTransactionNotFound)
	ErrStoreNotFound           = newErrorCode(ErrorCodeStoreNotFound)
	ErrRefundExceedLimitPerDay = newErrorCode(ErrorCodeRefundAmountExceedPerDay)
	ErrValidation              = newErrorCode(ErrorCodeValidationError)
)

type errorCode struct{ id string }

var (
	_ error     = (*errorCode)(nil)
	_ ErrorCode = (*errorCode)(nil)
)

func newErrorCode(id string) error {
	return &errorCode{id: id}
}

func (err errorCode) isCode(id string) bool {
	return err.id == id
}

func (err errorCode) Error() string {
	return fmt.Sprintf(errTemplate, err.id)
}

type ErrorCode interface {
	isCode(id string) bool
}

// Error :
type Error struct {
	code        string
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
	e.code = strings.ToUpper(strings.TrimSpace(gjson.GetBytes(respBytes, "error.code").String()))
	e.url = url
	e.rawResponse = respBytes
	e.rawRequest = reqBytes
	return e
}

func (e Error) isCode(errID string) bool {
	return e.code == errID
}

func (e Error) Is(err error) bool {
	v, ok := err.(ErrorCode)
	if ok {
		return v.isCode(e.code)
	}
	return e.Error() == err.Error()
}

// Error :
func (e Error) Error() string {
	return fmt.Sprintf(errTemplate, e.code)
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
