package rm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRmError(t *testing.T) {
	errc := newErrorCode(ErrorCodePaymentAlreadyRefunded)
	require.Implements(t, (*ErrorCode)(nil), errc)
	require.Implements(t, (*error)(nil), errc)
	require.Equal(t, "rm: PAYMENT_FULLY_REFUNDED", errc.Error())
	require.True(t, strings.HasPrefix(errc.Error(), "rm:"))
	require.True(t, errc.(ErrorCode).isCode(ErrorCodePaymentAlreadyRefunded))

	b, err := ioutil.ReadFile("./sample/validation_error.json")
	require.NoError(t, err)

	rmErr := newError("http://google.com", nil, b)
	require.Error(t, rmErr)
	require.Contains(t, fmt.Sprintf("%+v", rmErr), string(b))
	require.Implements(t, (*ErrorCode)(nil), rmErr)
	require.Implements(t, (*error)(nil), rmErr)
	require.True(t, strings.HasPrefix(rmErr.Error(), "rm:"))
	require.True(t, strings.HasPrefix(fmt.Sprint(rmErr), "rm:"))
	require.True(t, rmErr.isCode(ErrorCodeValidationError))
	require.True(t, rmErr.Is(ErrValidation))
	require.True(t, errors.Is(fmt.Errorf("wrap: %w", rmErr), ErrValidation))
	require.Equal(t, string(b), rmErr.Response())
	require.Equal(t, b, rmErr.ResponseBytes())
}
