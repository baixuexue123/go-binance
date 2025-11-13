package alpha

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/adshao/go-binance/v2/common"
)

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// OrderStatusType define order status type
type OrderStatusType string

// WalletType define wallet type
type WalletType string

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit OrderType = "LIMIT"

	TimeInForceTypeGTC TimeInForceType = "GTC"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"

	WalletTypeFunding WalletType = "FUNDING"
	WalletTypeAlpha   WalletType = "ALPHA"
)

// Endpoints
var (
	BaseAPIMainURL    = "https://api.binance.com"
	BaseAPITestnetURL = "https://testnet.binance.vision"
)

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned // if the 'timestamp' parameter is required
)

type params map[string]interface{}

// request define an API request
type request struct {
	method     string
	endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	secType    secType
	header     http.Header
	body       io.Reader
	fullURL    string
}

// addParam add param with key/value to query string
func (r *request) addParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string
func (r *request) setParams(m params) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

// setFormParam set param with key/value to request form body
func (r *request) setFormParam(key string, value interface{}) *request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setFormParams set params with key/values to request form body
func (r *request) setFormParams(m params) *request {
	for k, v := range m {
		r.setFormParam(k, v)
	}
	return r
}

func (r *request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return nil
}

// RequestOption define option type for request
type RequestOption func(*request)

// WithRecvWindow set recvWindow param for the request
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *request) {
		r.recvWindow = recvWindow
	}
}

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Logger     *log.Logger
	do         func(*http.Request) (*http.Response, error)
	TimeOffset int64 // time offset between server and client
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   BaseAPIMainURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		Logger: log.New(os.Stderr, "Binance-alpha ", log.LstdFlags),
	}
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Logger != nil {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from the user
	for _, opt := range opts {
		opt(r)
	}

	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}

	if r.body != nil {
		body = bytes.NewBuffer([]byte{})
		io.Copy(body, r.body)
	} else if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}

	if r.recvWindow > 0 {
		r.setParam("recvWindow", r.recvWindow)
		queryString = r.query.Encode()
	}

	if r.secType == secTypeSigned || r.secType == secTypeAPIKey {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.secType == secTypeSigned {
		queryString = r.query.Encode()
		timestamp := time.Now().UnixMilli() - c.TimeOffset
		if queryString != "" {
			queryString += "&"
		}
		queryString += fmt.Sprintf("timestamp=%d", timestamp)
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		sign, err := common.Hmac(c.SecretKey, raw)
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set("signature", *sign)
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}

	if queryString != "" {
		fullURL = fullURL + "?" + queryString
	}

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v\n", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v\n", res)
	c.debug("response body: %s\n", string(data))
	c.debug("response status code: %d\n", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s\n", e)
		}
		if !apiErr.IsValid() {
			apiErr.Response = data
		}
		return nil, apiErr
	}
	return data, nil
}

// SetApiEndpoint set api Endpoint
func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}

// SetTLSConfig set tls config
func (c *Client) SetTLSConfig(config *tls.Config) *Client {
	c.HTTPClient.Transport = &http.Transport{
		TLSClientConfig: config,
	}
	return c
}

// NewGetQuoteAssetsService init get quote assets service
func (c *Client) NewGetQuoteAssetsService() *GetQuoteAssetsService {
	return &GetQuoteAssetsService{c: c}
}

// NewGetTokenInfoService init get token info service
func (c *Client) NewGetTokenInfoService() *GetTokenInfoService {
	return &GetTokenInfoService{c: c}
}

// NewGetExchangeInfoService init get exchange info service
func (c *Client) NewGetExchangeInfoService() *GetExchangeInfoService {
	return &GetExchangeInfoService{c: c}
}

// NewGetCommissionFeeService init get commission fee service
func (c *Client) NewGetCommissionFeeService() *GetCommissionFeeService {
	return &GetCommissionFeeService{c: c}
}

// NewGetListenKeyService init get listen key service
func (c *Client) NewGetListenKeyService() *GetListenKeyService {
	return &GetListenKeyService{c: c}
}

// NewPlaceOrderService init place order service
func (c *Client) NewPlaceOrderService() *PlaceOrderService {
	return &PlaceOrderService{c: c}
}

// NewCancelOrderService init cancel order service
func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

// NewCancelAllOrdersService init cancel all orders service
func (c *Client) NewCancelAllOrdersService() *CancelAllOrdersService {
	return &CancelAllOrdersService{c: c}
}

// NewGetOpenOrdersService init get open orders service
func (c *Client) NewGetOpenOrdersService() *GetOpenOrdersService {
	return &GetOpenOrdersService{c: c}
}

// NewGetOrderHistoryService init get order history service
func (c *Client) NewGetOrderHistoryService() *GetOrderHistoryService {
	return &GetOrderHistoryService{c: c}
}

// NewGetOrderDetailService init get order detail service
func (c *Client) NewGetOrderDetailService() *GetOrderDetailService {
	return &GetOrderDetailService{c: c}
}

// NewGetUserTradesService init get user trades service
func (c *Client) NewGetUserTradesService() *GetUserTradesService {
	return &GetUserTradesService{c: c}
}

// NewGetKlinesService init get klines service
func (c *Client) NewGetKlinesService() *GetKlinesService {
	return &GetKlinesService{c: c}
}

// NewGetTickerService init get ticker service
func (c *Client) NewGetTickerService() *GetTickerService {
	return &GetTickerService{c: c}
}

// NewGetTickerPriceService init get ticker price service
func (c *Client) NewGetTickerPriceService() *GetTickerPriceService {
	return &GetTickerPriceService{c: c}
}

// NewGetAggTradesService init get agg trades service
func (c *Client) NewGetAggTradesService() *GetAggTradesService {
	return &GetAggTradesService{c: c}
}

// NewGetBookTickerService init get book ticker service
func (c *Client) NewGetBookTickerService() *GetBookTickerService {
	return &GetBookTickerService{c: c}
}

// NewGetDepthService init get depth service
func (c *Client) NewGetDepthService() *GetDepthService {
	return &GetDepthService{c: c}
}

// NewGetAlphaAssetService init get alpha asset service
func (c *Client) NewGetAlphaAssetService() *GetAlphaAssetService {
	return &GetAlphaAssetService{c: c}
}

// NewGetTokenMappingService init get token mapping service
func (c *Client) NewGetTokenMappingService() *GetTokenMappingService {
	return &GetTokenMappingService{c: c}
}

// NewAlphaWithdrawService init alpha withdraw service
func (c *Client) NewAlphaWithdrawService() *AlphaWithdrawService {
	return &AlphaWithdrawService{c: c}
}

// NewGetAlphaWithdrawHistoryService init get alpha withdraw history service
func (c *Client) NewGetAlphaWithdrawHistoryService() *GetAlphaWithdrawHistoryService {
	return &GetAlphaWithdrawHistoryService{c: c}
}

// NewGetAlphaDepositHistoryService init get alpha deposit history service
func (c *Client) NewGetAlphaDepositHistoryService() *GetAlphaDepositHistoryService {
	return &GetAlphaDepositHistoryService{c: c}
}

// NewGetAlphaDepositAddressService init get alpha deposit address service
func (c *Client) NewGetAlphaDepositAddressService() *GetAlphaDepositAddressService {
	return &GetAlphaDepositAddressService{c: c}
}
