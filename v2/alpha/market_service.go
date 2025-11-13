package alpha

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetQuoteAssetsService fetch quote assets supported on alpha market
type GetQuoteAssetsService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetQuoteAssetsService) RecvWindow(recvWindow int64) *GetQuoteAssetsService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetQuoteAssetsService) Do(ctx context.Context, opts ...RequestOption) ([]string, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/get-from-asset",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []string
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetTokenInfoService fetch alpha token & network info
type GetTokenInfoService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetTokenInfoService) RecvWindow(recvWindow int64) *GetTokenInfoService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetTokenInfoService) Do(ctx context.Context, opts ...RequestOption) ([]TokenConfig, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/alpha/config/getall",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []TokenConfig
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetExchangeInfoService fetch all exchange info supported on alpha market
type GetExchangeInfoService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetExchangeInfoService) RecvWindow(recvWindow int64) *GetExchangeInfoService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (*ExchangeInfo, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/get-exchange-info",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ExchangeInfo
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCommissionFeeService fetch all commission fee info by symbol
type GetCommissionFeeService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetCommissionFeeService) Symbol(symbol string) *GetCommissionFeeService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *GetCommissionFeeService) RecvWindow(recvWindow int64) *GetCommissionFeeService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetCommissionFeeService) Do(ctx context.Context, opts ...RequestOption) (*CommissionFee, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/get-fee-rate",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res CommissionFee
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetKlinesService fetch kline data supported on alpha market
type GetKlinesService struct {
	c          *Client
	symbol     string
	interval   string
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *GetKlinesService) Symbol(symbol string) *GetKlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *GetKlinesService) Interval(interval string) *GetKlinesService {
	s.interval = interval
	return s
}

// StartTime set startTime
func (s *GetKlinesService) StartTime(startTime int64) *GetKlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetKlinesService) EndTime(endTime int64) *GetKlinesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetKlinesService) Limit(limit int) *GetKlinesService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *GetKlinesService) RecvWindow(recvWindow int64) *GetKlinesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetKlinesService) Do(ctx context.Context, opts ...RequestOption) ([]Kline, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/klines",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("interval", s.interval)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var rawKlines [][]interface{}
	if err = json.Unmarshal(data, &rawKlines); err != nil {
		return nil, err
	}
	klines := make([]Kline, 0, len(rawKlines))
	for _, k := range rawKlines {
		if len(k) < 12 {
			continue
		}
		kl := Kline{}
		if v, ok := k[0].(float64); ok {
			kl.OpenTime = int64(v)
		}
		if v, ok := k[1].(string); ok {
			kl.Open = v
		}
		if v, ok := k[2].(string); ok {
			kl.High = v
		}
		if v, ok := k[3].(string); ok {
			kl.Low = v
		}
		if v, ok := k[4].(string); ok {
			kl.Close = v
		}
		if v, ok := k[5].(string); ok {
			kl.Volume = v
		}
		if v, ok := k[6].(float64); ok {
			kl.CloseTime = int64(v)
		}
		if v, ok := k[7].(string); ok {
			kl.QuoteAssetVolume = v
		}
		if v, ok := k[8].(float64); ok {
			kl.NumberOfTrades = int(v)
		}
		if v, ok := k[9].(string); ok {
			kl.TakerBuyBaseAssetVolume = v
		}
		if v, ok := k[10].(string); ok {
			kl.TakerBuyQuoteAssetVolume = v
		}
		if v, ok := k[11].(string); ok {
			kl.Ignore = v
		}
		klines = append(klines, kl)
	}
	return klines, nil
}

// GetTickerService fetch 24hr ticker price change statistics
type GetTickerService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetTickerService) Symbol(symbol string) *GetTickerService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *GetTickerService) RecvWindow(recvWindow int64) *GetTickerService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetTickerService) Do(ctx context.Context, opts ...RequestOption) (*Ticker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/ticker",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res Ticker
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTickerPriceService fetch symbol price ticker
type GetTickerPriceService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetTickerPriceService) Symbol(symbol string) *GetTickerPriceService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *GetTickerPriceService) RecvWindow(recvWindow int64) *GetTickerPriceService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetTickerPriceService) Do(ctx context.Context, opts ...RequestOption) (*TickerPrice, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/ticker-price",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res TickerPrice
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAggTradesService fetch aggregate trade list
type GetAggTradesService struct {
	c          *Client
	symbol     string
	fromID     *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *GetAggTradesService) Symbol(symbol string) *GetAggTradesService {
	s.symbol = symbol
	return s
}

// FromID set fromID
func (s *GetAggTradesService) FromID(fromID int64) *GetAggTradesService {
	s.fromID = &fromID
	return s
}

// StartTime set startTime
func (s *GetAggTradesService) StartTime(startTime int64) *GetAggTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetAggTradesService) EndTime(endTime int64) *GetAggTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAggTradesService) Limit(limit int) *GetAggTradesService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *GetAggTradesService) RecvWindow(recvWindow int64) *GetAggTradesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetAggTradesService) Do(ctx context.Context, opts ...RequestOption) ([]AggTrade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/agg-trades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []AggTrade
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetBookTickerService fetch symbol order book ticker
type GetBookTickerService struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetBookTickerService) Symbol(symbol string) *GetBookTickerService {
	s.symbol = symbol
	return s
}

// RecvWindow set recvWindow
func (s *GetBookTickerService) RecvWindow(recvWindow int64) *GetBookTickerService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetBookTickerService) Do(ctx context.Context, opts ...RequestOption) (*BookTicker, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/book-ticker",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res BookTicker
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetDepthService fetch order book depth
type GetDepthService struct {
	c          *Client
	symbol     string
	limit      *int
	recvWindow *int64
}

// Symbol set symbol
func (s *GetDepthService) Symbol(symbol string) *GetDepthService {
	s.symbol = symbol
	return s
}

// Limit set limit
func (s *GetDepthService) Limit(limit int) *GetDepthService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *GetDepthService) RecvWindow(recvWindow int64) *GetDepthService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetDepthService) Do(ctx context.Context, opts ...RequestOption) (*Depth, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/market/depth",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res Depth
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
