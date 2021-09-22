package rm

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRmError(t *testing.T) {
	b, err := ioutil.ReadFile("./sample/validation_error.json")
	require.NoError(t, err)

	rmErr := newError("http://google.com", nil, b)
	require.Error(t, rmErr)
	require.Equal(t, `URL : http://google.com
Request Body :

Response Body :
`+string(b), fmt.Sprintf("%+v", rmErr))
	require.True(t, strings.HasPrefix(rmErr.Error(), "rm:"))
	require.True(t, strings.HasPrefix(fmt.Sprint(rmErr), "rm:"))
	require.Equal(t, string(b), rmErr.Response())
	require.Equal(t, b, rmErr.ResponseBytes())
}
