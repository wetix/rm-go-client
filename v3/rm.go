package rm

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/clbanning/mxj"
	"github.com/dchest/uniuri"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/valyala/bytebufferpool"
	"golang.org/x/oauth2"
)

const (
	ResponseSuccess = "SUCCESS"
)

var (
	noopTracer = &opentracing.NoopTracer{}
)

type Config struct {
	ClientID     string
	ClientSecret string
	PrivateKey   []byte
	PublicKey    []byte
	StoreID      string
	Sandbox      bool
	TokenSource  oauth2.TokenSource
}

// Client :
type Client struct {
	mu            sync.Mutex
	clientID      string
	clientSecret  string
	sandbox       bool
	oauthEndpoint string
	openEndpoint  string
	token         *oauth2.Token
	pk            *rsa.PrivateKey
	pub           []byte
	oauth2        oauth2.TokenSource
	storeID       string
}

// NewClient :
func NewClient(cfg Config) *Client {
	var (
		c   = new(Client)
		err error
	)
	c.clientID = cfg.ClientID
	c.clientSecret = cfg.ClientSecret
	c.oauthEndpoint = "https://oauth.revenuemonster.my"
	c.openEndpoint = "https://open.revenuemonster.my"
	if cfg.Sandbox {
		c.oauthEndpoint = "https://sb-oauth.revenuemonster.my"
		c.openEndpoint = "https://sb-open.revenuemonster.my"
	}

	block, _ := pem.Decode(cfg.PrivateKey)
	if block == nil {
		panic("invalid format of private key")
	}

	c.pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	c.pub = cfg.PublicKey
	if cfg.TokenSource != nil {
		c.oauth2 = cfg.TokenSource
	} else {
		c.oauth2 = c
	}

	// if store id is empty, get store id
	if cfg.StoreID == "" {
		resp, err := c.GetStores(context.Background())
		if err != nil {
			panic(err)
		}
		c.storeID = resp.Items[0].ID
	}
	return c
}

func (c *Client) maybeStartSpanFromContext(ctx context.Context, operationName string) opentracing.Span {
	var span opentracing.Span
	if sp := opentracing.SpanFromContext(ctx); sp != nil {
		span, _ = opentracing.StartSpanFromContext(ctx, operationName)
	} else {
		span = noopTracer.StartSpan(operationName)
	}
	return span
}

func (c *Client) do(
	ctx context.Context,
	operationName string,
	method string,
	endpoint string,
	src interface{},
	dest interface{},
) error {
	var (
		req    = new(http.Request)
		b      = make([]byte, 0)
		b64Str string
		sign   string
		err    error
	)

	span := c.maybeStartSpanFromContext(ctx, operationName)
	defer span.Finish()

	defer func() {
		if err != nil {
			ext.LogError(span, err)
		}
	}()

	if src != nil {
		b, err = json.Marshal(src)
		if err != nil {
			return err
		}
	}

	method = strings.TrimSpace(strings.ToLower(method))
	reqUrl, _ := url.Parse(endpoint)
	req.Method = strings.ToUpper(method)
	req.URL = reqUrl

	ext.HTTPUrl.Set(span, endpoint)
	ext.HTTPMethod.Set(span, method)
	ext.Component.Set(span, "rm-go-client")

	span.LogFields(
		log.String("http.request.body", string(b)),
	)

	if len(b) > 0 &&
		!bytes.Equal(b, []byte(`null`)) &&
		!bytes.Equal(b, []byte(`{}`)) {

		var (
			buf = new(bytes.Buffer)
			m   mxj.Map
			js  []byte
		)

		m, err = mxj.NewMapJson(b)
		if err != nil {
			return err
		}

		js, err = m.Json(true)
		if err != nil {
			return err
		}

		err = json.Compact(buf, js)
		if err != nil {
			return err
		}

		req.Body = ioutil.NopCloser(buf)
		b64Str = base64.StdEncoding.EncodeToString(buf.Bytes())
	}

	var tkn *oauth2.Token
	tkn, err = c.oauth2.Token()
	if err != nil {
		return err
	}

	data := []string{}
	randomStr := uniuri.NewLen(25)
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	if b64Str != "" {
		data = append(data, "data="+b64Str)
	}
	data = append(data, "method="+method)
	data = append(data, "nonceStr="+randomStr)
	data = append(data, "requestUrl="+endpoint)
	data = append(data, "signType=sha256")
	data = append(data, "timestamp="+ts)

	sign, err = signData(crypto.SHA256, data, c.pk)
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"Accept":        {"application/json"},
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + tkn.AccessToken},
		"X-Nonce-Str":   {randomStr},
		"X-Signature":   {"sha256 " + sign},
		"X-Timestamp":   {ts},
	}

	var res *http.Response
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	ext.HTTPStatusCode.Set(span, uint16(res.StatusCode))

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	span.LogFields(
		log.String("http.response.body", string(b)),
	)

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		return newError(b)
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return err
	}
	return nil
}

func signData(h crypto.Hash, data []string, pk *rsa.PrivateKey) (string, error) {
	hash, err := signPKCS1v15(h, data, pk)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}

func signPKCS1v15(hash crypto.Hash, data []string, pk *rsa.PrivateKey) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	for idx := range data {
		if idx > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(data[idx])
	}

	h := hash.New()
	if _, err := h.Write(buf.Bytes()); err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, pk, hash, h.Sum(nil))
}
