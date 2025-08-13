package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListAllConvertPairsService
// https://binance-docs.github.io/apidocs/spot/en/#list-all-convert-pairs
type ListAllConvertPairsService struct {
	c         *Client
	fromAsset string
	toAsset   string
}

func (s *ListAllConvertPairsService) FromAsset(v string) *ListAllConvertPairsService {
	s.fromAsset = v
	return s
}

func (s *ListAllConvertPairsService) ToAsset(v string) *ListAllConvertPairsService {
	s.toAsset = v
	return s
}

func (s *ListAllConvertPairsService) Do(ctx context.Context, opts ...RequestOption) ([]ConvertPair, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/exchangeInfo",
		secType:  secTypeSigned,
	}
	r.setParam("fromAsset", s.fromAsset)
	r.setParam("toAsset", s.toAsset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []ConvertPair
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type ConvertPair struct {
	FromAsset          string `json:"fromAsset"`
	ToAsset            string `json:"toAsset"`
	FromAssetMinAmount string `json:"fromAssetMinAmount"`
	FromAssetMaxAmount string `json:"fromAssetMaxAmount"`
	ToAssetMinAmount   string `json:"toAssetMinAmount"`
	ToAssetMaxAmount   string `json:"toAssetMaxAmount"`
}

// ConvertAssetInfoService
// https://binance-docs.github.io/apidocs/spot/en/#query-order-quantity-precision-per-asset-user_data
type ConvertAssetInfoService struct {
	c *Client
}

func (s *ConvertAssetInfoService) Do(ctx context.Context, opts ...RequestOption) ([]ConvertAssetInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/assetInfo",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []ConvertAssetInfo
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type ConvertAssetInfo struct {
	Asset    string `json:"asset"`
	Fraction int    `json:"fraction"`
}

// ConvertGetQuoteService
// https://binance-docs.github.io/apidocs/spot/en/#send-quote-request-user_data
type ConvertGetQuoteService struct {
	c          *Client
	fromAsset  string
	toAsset    string
	fromAmount string
	toAmount   string
	walletType string // SPOT or FUNDING. Default is SPOT
	validTime  string //10s, 30s, 1m, 2m, default 10s
}

func (s *ConvertGetQuoteService) FromAsset(v string) *ConvertGetQuoteService {
	s.fromAsset = v
	return s
}

func (s *ConvertGetQuoteService) ToAsset(v string) *ConvertGetQuoteService {
	s.toAsset = v
	return s
}

func (s *ConvertGetQuoteService) FromAmount(v string) *ConvertGetQuoteService {
	s.fromAmount = v
	return s
}

func (s *ConvertGetQuoteService) ToAmount(v string) *ConvertGetQuoteService {
	s.toAmount = v
	return s
}

func (s *ConvertGetQuoteService) WalletType(v string) *ConvertGetQuoteService {
	s.walletType = v
	return s
}

func (s *ConvertGetQuoteService) ValidTime(v string) *ConvertGetQuoteService {
	s.validTime = v
	return s
}

func (s *ConvertGetQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertGetQuoteResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/convert/getQuote",
		secType:  secTypeSigned,
	}
	r.setParam("fromAsset", s.fromAsset)
	r.setParam("toAsset", s.toAsset)
	if s.fromAmount != "" {
		r.setParam("fromAmount", s.fromAmount)
	}
	if s.toAmount != "" {
		r.setParam("toAmount", s.toAmount)
	}
	if s.walletType != "" {
		r.setParam("walletType", s.walletType)
	}
	if s.validTime != "" {
		r.setParam("validTime", s.validTime)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ConvertGetQuoteResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type ConvertGetQuoteResponse struct {
	QuoteID        string `json:"quoteId"`
	Ratio          string `json:"ratio"`
	InverseRatio   string `json:"inverseRatio"`
	ValidTimestamp int64  `json:"validTimestamp"`
	ToAmount       string `json:"toAmount"`
	FromAmount     string `json:"fromAmount"`
}

// ConvertAcceptQuoteService
// https://binance-docs.github.io/apidocs/spot/en/#accept-quote-trade
type ConvertAcceptQuoteService struct {
	c       *Client
	quoteId string
}

func (s *ConvertAcceptQuoteService) QuoteId(v string) *ConvertAcceptQuoteService {
	s.quoteId = v
	return s
}

func (s *ConvertAcceptQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertAcceptQuoteResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/convert/acceptQuote",
		secType:  secTypeSigned,
	}
	r.setParam("quoteId", s.quoteId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ConvertAcceptQuoteResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type ConvertAcceptQuoteResponse struct {
	OrderID     string `json:"orderId"`
	CreateTime  int64  `json:"createTime"`
	OrderStatus string `json:"orderStatus"`
}

// ConvertOrderStatusService
// https://binance-docs.github.io/apidocs/spot/en/#order-status-user_data
type ConvertOrderStatusService struct {
	c       *Client
	orderId string
	quoteId string
}

func (s *ConvertOrderStatusService) QuoteId(v string) *ConvertOrderStatusService {
	s.quoteId = v
	return s
}

func (s *ConvertOrderStatusService) Do(ctx context.Context, opts ...RequestOption) (*ConvertOrderDetail, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/orderStatus",
		secType:  secTypeSigned,
	}
	if s.orderId != "" {
		r.setParam("orderId", s.orderId)
	}
	if s.quoteId != "" {
		r.setParam("quoteId", s.quoteId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ConvertOrderDetail
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type ConvertOrderDetail struct {
	OrderID      int64  `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int64  `json:"createTime"`
}

type ConvertTradeHistoryService struct {
	c         *Client
	startTime int64
	endTime   int64
	limit     *int32
}

// StartTime set startTime
func (s *ConvertTradeHistoryService) StartTime(startTime int64) *ConvertTradeHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *ConvertTradeHistoryService) EndTime(endTime int64) *ConvertTradeHistoryService {
	s.endTime = endTime
	return s
}

// Limit set limit
func (s *ConvertTradeHistoryService) Limit(limit int32) *ConvertTradeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ConvertTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (*ConvertTradeHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/tradeFlow",
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertTradeHistory{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertTradeHistory define the convert trade history
type ConvertTradeHistory struct {
	List      []ConvertTradeHistoryItem `json:"list"`
	StartTime int64                     `json:"startTime"`
	EndTime   int64                     `json:"endTime"`
	Limit     int32                     `json:"limit"`
	MoreData  bool                      `json:"moreData"`
}

// ConvertTradeHistoryItem define a convert trade history item
type ConvertTradeHistoryItem struct {
	QuoteId      string `json:"quoteId"`
	OrderId      int64  `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int64  `json:"createTime"`
}

// ConvertExchangeInfoService create a new convert exchange info service
type ConvertExchangeInfoService struct {
	c         *Client
	fromAsset *string
	toAsset   *string
}

// FromAsset set fromAsset
func (s *ConvertExchangeInfoService) FromAsset(fromAsset string) *ConvertExchangeInfoService {
	s.fromAsset = &fromAsset
	return s
}

// ToAsset set toAsset
func (s *ConvertExchangeInfoService) ToAsset(toAsset string) *ConvertExchangeInfoService {
	s.toAsset = &toAsset
	return s
}

// Do send request
func (s *ConvertExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) ([]*ConvertExchangeInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/exchangeInfo",
		secType:  secTypeSigned,
	}
	if s.fromAsset != nil {
		r.setParam("fromAsset", *s.fromAsset)
	}
	if s.toAsset != nil {
		r.setParam("toAsset", *s.toAsset)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*ConvertExchangeInfo
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// ConvertExchangeInfo define the convert exchange info
type ConvertExchangeInfo struct {
	FromAsset          string `json:"fromAsset"`
	ToAsset            string `json:"toAsset"`
	FromAssetMinAmount string `json:"fromAssetMinAmount"`
	FromAssetMaxAmount string `json:"fromAssetMaxAmount"`
	ToAssetMinAmount   string `json:"toAssetMinAmount"`
	ToAssetMaxAmount   string `json:"toAssetMaxAmount"`
}

// ConvertGetQuoteService create a new convert quote service
type ConvertGetQuoteService struct {
	c          *Client
	fromAsset  string
	toAsset    string
	fromAmount *string
	toAmount   *string
	walletType *string // SPOT, FUNDING
	validTime  *string // 10s, 30s, 1m, 2m, default 10s
}

// ConvertAssetInfoService create a new convert asset info service
type ConvertAssetInfoService struct {
	c *Client
}

// Do send request
func (s *ConvertAssetInfoService) Do(ctx context.Context, opts ...RequestOption) ([]*ConvertAssetInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/assetInfo",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*ConvertAssetInfo
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// ConvertAssetInfo define the convert asset info
type ConvertAssetInfo struct {
	Asset    string `json:"asset"`
	Fraction int    `json:"fraction"`
}

// FromAsset set fromAsset
func (s *ConvertGetQuoteService) FromAsset(fromAsset string) *ConvertGetQuoteService {
	s.fromAsset = fromAsset
	return s
}

// ToAsset set fromAsset
func (s *ConvertGetQuoteService) ToAsset(toAsset string) *ConvertGetQuoteService {
	s.toAsset = toAsset
	return s
}

// FromAmount set fromAmount
func (s *ConvertGetQuoteService) FromAmount(fromAmount string) *ConvertGetQuoteService {
	s.fromAmount = &fromAmount
	return s
}

// ToAmount set toAmount
func (s *ConvertGetQuoteService) ToAmount(toAmount string) *ConvertGetQuoteService {
	s.toAmount = &toAmount
	return s
}

// WalletType set walletType
// SPOT or FUNDING. Default is SPOT
func (s *ConvertGetQuoteService) WalletType(walletType string) *ConvertGetQuoteService {
	s.walletType = &walletType
	return s
}

// ValidTime set validTime
// 10s, 30s, 1m, 2m, default 10s
func (s *ConvertGetQuoteService) ValidTime(validTime string) *ConvertGetQuoteService {
	s.validTime = &validTime
	return s
}

// Do send request
func (s *ConvertGetQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertQuote, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/convert/getQuote",
		secType:  secTypeSigned,
	}

	r.setParam("fromAsset", s.fromAsset)
	r.setParam("toAsset", s.toAsset)

	if s.fromAmount != nil {
		r.setParam("fromAmount", *s.fromAmount)
	}
	if s.toAmount != nil {
		r.setParam("toAmount", *s.toAmount)
	}
	if s.walletType != nil {
		r.setParam("walletType", *s.walletType)
	}
	if s.validTime != nil {
		r.setParam("validTime", *s.validTime)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ConvertQuote
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertQuote define the convert quote
type ConvertQuote struct {
	QuoteId      string `json:"quoteId"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	ValidTime    int64  `json:"validTime"`
	ToAmount     string `json:"toAmount"`
	FromAmount   string `json:"fromAmount"`
}

type ConvertAcceptQuoteService struct {
	c       *Client
	quoteId string
}

// QuoteId set quoteId
func (s *ConvertAcceptQuoteService) QuoteId(quoteId string) *ConvertAcceptQuoteService {
	s.quoteId = quoteId
	return s
}

// Do send request
func (s *ConvertAcceptQuoteService) Do(ctx context.Context, opts ...RequestOption) (*ConvertAcceptQuote, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/convert/acceptQuote",
		secType:  secTypeSigned,
	}
	r.setParam("quoteId", s.quoteId)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ConvertAcceptQuote
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertAcceptQuote define the convert accept quote
type ConvertAcceptQuote struct {
	OrderId     string `json:"orderId"`
	CreateTime  int64  `json:"createTime"`
	OrderStatus string `json:"orderStatus"`
}

// ConvertOrderStatusService check order status
type ConvertOrderStatusService struct {
	c       *Client
	orderId *string
	quoteId *string
}

// OrderId set orderId
func (s *ConvertOrderStatusService) OrderId(orderId string) *ConvertOrderStatusService {
	s.orderId = &orderId
	return s
}

// QuoteId set quoteId
func (s *ConvertOrderStatusService) QuoteId(quoteId string) *ConvertOrderStatusService {
	s.quoteId = &quoteId
	return s
}

// Do send request
func (s *ConvertOrderStatusService) Do(ctx context.Context, opts ...RequestOption) (*ConvertOrderStatus, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/orderStatus",
		secType:  secTypeSigned,
	}
	if s.orderId != nil {
		r.setParam("orderId", *s.orderId)
	}
	if s.quoteId != nil {
		r.setParam("quoteId", *s.quoteId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertOrderStatus{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ConvertOrderStatus define the convert order status
type ConvertOrderStatus struct {
	OrderId      int64  `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int64  `json:"createTime"`
}
