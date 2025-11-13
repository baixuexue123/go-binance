package alpha

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetListenKeyService get listen key to subscribe websocket
type GetListenKeyService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetListenKeyService) RecvWindow(recvWindow int64) *GetListenKeyService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetListenKeyService) Do(ctx context.Context, opts ...RequestOption) (*ListenKey, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/alpha-trade/get-listen-key",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res ListenKey
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PlaceOrderService place alpha order
type PlaceOrderService struct {
	c             *Client
	baseAsset     string
	quoteAsset    string
	side          SideType
	quantity      string
	price         string
	clientOrderID *string
	walletType    *WalletType
	recvWindow    *int64
}

// BaseAsset set baseAsset
func (s *PlaceOrderService) BaseAsset(baseAsset string) *PlaceOrderService {
	s.baseAsset = baseAsset
	return s
}

// QuoteAsset set quoteAsset
func (s *PlaceOrderService) QuoteAsset(quoteAsset string) *PlaceOrderService {
	s.quoteAsset = quoteAsset
	return s
}

// Side set side
func (s *PlaceOrderService) Side(side SideType) *PlaceOrderService {
	s.side = side
	return s
}

// Quantity set quantity
func (s *PlaceOrderService) Quantity(quantity string) *PlaceOrderService {
	s.quantity = quantity
	return s
}

// Price set price
func (s *PlaceOrderService) Price(price string) *PlaceOrderService {
	s.price = price
	return s
}

// ClientOrderID set clientOrderId
func (s *PlaceOrderService) ClientOrderID(clientOrderID string) *PlaceOrderService {
	s.clientOrderID = &clientOrderID
	return s
}

// WalletType set walletType
func (s *PlaceOrderService) WalletType(walletType WalletType) *PlaceOrderService {
	s.walletType = &walletType
	return s
}

// RecvWindow set recvWindow
func (s *PlaceOrderService) RecvWindow(recvWindow int64) *PlaceOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *PlaceOrderService) Do(ctx context.Context, opts ...RequestOption) (*PlaceOrderResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/alpha-trade/order/place",
		secType:  secTypeSigned,
	}
	r.setFormParam("baseAsset", s.baseAsset)
	r.setFormParam("quoteAsset", s.quoteAsset)
	r.setFormParam("side", s.side)
	r.setFormParam("quantity", s.quantity)
	r.setFormParam("price", s.price)
	if s.clientOrderID != nil {
		r.setFormParam("clientOrderId", *s.clientOrderID)
	}
	if s.walletType != nil {
		r.setFormParam("walletType", *s.walletType)
	}
	if s.recvWindow != nil {
		r.setFormParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res PlaceOrderResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelOrderService cancel alpha order
type CancelOrderService struct {
	c          *Client
	symbol     string
	orderID    string
	recvWindow *int64
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderId
func (s *CancelOrderService) OrderID(orderID string) *CancelOrderService {
	s.orderID = orderID
	return s
}

// RecvWindow set recvWindow
func (s *CancelOrderService) RecvWindow(recvWindow int64) *CancelOrderService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (*CancelOrderResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/alpha-trade/order/cancel",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	r.setFormParam("orderId", s.orderID)
	if s.recvWindow != nil {
		r.setFormParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res CancelOrderResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelAllOrdersService cancel all alpha orders
type CancelAllOrdersService struct {
	c          *Client
	symbol     *string
	baseAsset  *string
	recvWindow *int64
}

// Symbol set symbol
func (s *CancelAllOrdersService) Symbol(symbol string) *CancelAllOrdersService {
	s.symbol = &symbol
	return s
}

// BaseAsset set baseAsset
func (s *CancelAllOrdersService) BaseAsset(baseAsset string) *CancelAllOrdersService {
	s.baseAsset = &baseAsset
	return s
}

// RecvWindow set recvWindow
func (s *CancelAllOrdersService) RecvWindow(recvWindow int64) *CancelAllOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *CancelAllOrdersService) Do(ctx context.Context, opts ...RequestOption) (*CancelAllOrdersResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/alpha-trade/order/cancel-all",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setFormParam("symbol", *s.symbol)
	}
	if s.baseAsset != nil {
		r.setFormParam("baseAsset", *s.baseAsset)
	}
	if s.recvWindow != nil {
		r.setFormParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res CancelAllOrdersResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetOpenOrdersService get open orders
type GetOpenOrdersService struct {
	c          *Client
	symbol     *string
	side       *SideType
	recvWindow *int64
}

// Symbol set symbol
func (s *GetOpenOrdersService) Symbol(symbol string) *GetOpenOrdersService {
	s.symbol = &symbol
	return s
}

// Side set side
func (s *GetOpenOrdersService) Side(side SideType) *GetOpenOrdersService {
	s.side = &side
	return s
}

// RecvWindow set recvWindow
func (s *GetOpenOrdersService) RecvWindow(recvWindow int64) *GetOpenOrdersService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) ([]Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/order/get-open-order",
		secType:  secTypeSigned,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	if s.side != nil {
		r.setParam("side", *s.side)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []Order
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetOrderHistoryService get order history
type GetOrderHistoryService struct {
	c           *Client
	baseAsset   *string
	side        *SideType
	orderStatus *string
	startTime   *int64
	endTime     *int64
	limit       *int64
	pageID      *int64
	recvWindow  *int64
}

// BaseAsset set baseAsset
func (s *GetOrderHistoryService) BaseAsset(baseAsset string) *GetOrderHistoryService {
	s.baseAsset = &baseAsset
	return s
}

// Side set side
func (s *GetOrderHistoryService) Side(side SideType) *GetOrderHistoryService {
	s.side = &side
	return s
}

// OrderStatus set orderStatus
func (s *GetOrderHistoryService) OrderStatus(orderStatus string) *GetOrderHistoryService {
	s.orderStatus = &orderStatus
	return s
}

// StartTime set startTime
func (s *GetOrderHistoryService) StartTime(startTime int64) *GetOrderHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetOrderHistoryService) EndTime(endTime int64) *GetOrderHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetOrderHistoryService) Limit(limit int64) *GetOrderHistoryService {
	s.limit = &limit
	return s
}

// PageID set pageId
func (s *GetOrderHistoryService) PageID(pageID int64) *GetOrderHistoryService {
	s.pageID = &pageID
	return s
}

// RecvWindow set recvWindow
func (s *GetOrderHistoryService) RecvWindow(recvWindow int64) *GetOrderHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetOrderHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/order/get-order-history",
		secType:  secTypeSigned,
	}
	if s.baseAsset != nil {
		r.setParam("baseAsset", *s.baseAsset)
	}
	if s.side != nil {
		r.setParam("side", *s.side)
	}
	if s.orderStatus != nil {
		r.setParam("orderStatus", *s.orderStatus)
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
	if s.pageID != nil {
		r.setParam("pageId", *s.pageID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []Order
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetOrderDetailService get order detail
type GetOrderDetailService struct {
	c          *Client
	symbol     string
	orderID    string
	recvWindow *int64
}

// Symbol set symbol
func (s *GetOrderDetailService) Symbol(symbol string) *GetOrderDetailService {
	s.symbol = symbol
	return s
}

// OrderID set orderId
func (s *GetOrderDetailService) OrderID(orderID string) *GetOrderDetailService {
	s.orderID = orderID
	return s
}

// RecvWindow set recvWindow
func (s *GetOrderDetailService) RecvWindow(recvWindow int64) *GetOrderDetailService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetOrderDetailService) Do(ctx context.Context, opts ...RequestOption) (*Order, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/order/get-order-detail",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	r.setParam("orderId", s.orderID)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res Order
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserTradesService get user trades
type GetUserTradesService struct {
	c          *Client
	baseAsset  *string
	side       *SideType
	orderID    *string
	startTime  *int64
	endTime    *int64
	limit      *int64
	pageID     *int64
	recvWindow *int64
}

// BaseAsset set baseAsset
func (s *GetUserTradesService) BaseAsset(baseAsset string) *GetUserTradesService {
	s.baseAsset = &baseAsset
	return s
}

// Side set side
func (s *GetUserTradesService) Side(side SideType) *GetUserTradesService {
	s.side = &side
	return s
}

// OrderID set orderId
func (s *GetUserTradesService) OrderID(orderID string) *GetUserTradesService {
	s.orderID = &orderID
	return s
}

// StartTime set startTime
func (s *GetUserTradesService) StartTime(startTime int64) *GetUserTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetUserTradesService) EndTime(endTime int64) *GetUserTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetUserTradesService) Limit(limit int64) *GetUserTradesService {
	s.limit = &limit
	return s
}

// PageID set pageId
func (s *GetUserTradesService) PageID(pageID int64) *GetUserTradesService {
	s.pageID = &pageID
	return s
}

// RecvWindow set recvWindow
func (s *GetUserTradesService) RecvWindow(recvWindow int64) *GetUserTradesService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUserTradesService) Do(ctx context.Context, opts ...RequestOption) ([]Trade, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/order/get-user-trades",
		secType:  secTypeSigned,
	}
	if s.baseAsset != nil {
		r.setParam("baseAsset", *s.baseAsset)
	}
	if s.side != nil {
		r.setParam("side", *s.side)
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
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
	if s.pageID != nil {
		r.setParam("pageId", *s.pageID)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []Trade
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
