package alpha

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAlphaAssetService fetch alpha wallet asset balance
type GetAlphaAssetService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetAlphaAssetService) RecvWindow(recvWindow int64) *GetAlphaAssetService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetAlphaAssetService) Do(ctx context.Context, opts ...RequestOption) ([]AlphaAsset, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/get-alpha-asset",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []AlphaAsset
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetTokenMappingService fetch alphaId & CA mapping info
type GetTokenMappingService struct {
	c          *Client
	recvWindow *int64
}

// RecvWindow set recvWindow
func (s *GetTokenMappingService) RecvWindow(recvWindow int64) *GetTokenMappingService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetTokenMappingService) Do(ctx context.Context, opts ...RequestOption) ([]TokenMapping, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/alpha-trade/token/all/list",
		secType:  secTypeSigned,
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []TokenMapping
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// AlphaWithdrawService submit a withdrawal request
type AlphaWithdrawService struct {
	c               *Client
	network         string
	alphaID         string
	contractAddress string
	address         string
	addressTag      *string
	amount          string
	clientOrderID   *string
	recvWindow      *int64
}

// Network set network
func (s *AlphaWithdrawService) Network(network string) *AlphaWithdrawService {
	s.network = network
	return s
}

// AlphaID set alphaId
func (s *AlphaWithdrawService) AlphaID(alphaID string) *AlphaWithdrawService {
	s.alphaID = alphaID
	return s
}

// ContractAddress set contractAddress
func (s *AlphaWithdrawService) ContractAddress(contractAddress string) *AlphaWithdrawService {
	s.contractAddress = contractAddress
	return s
}

// Address set address
func (s *AlphaWithdrawService) Address(address string) *AlphaWithdrawService {
	s.address = address
	return s
}

// AddressTag set addressTag
func (s *AlphaWithdrawService) AddressTag(addressTag string) *AlphaWithdrawService {
	s.addressTag = &addressTag
	return s
}

// Amount set amount
func (s *AlphaWithdrawService) Amount(amount string) *AlphaWithdrawService {
	s.amount = amount
	return s
}

// ClientOrderID set clientOrderId
func (s *AlphaWithdrawService) ClientOrderID(clientOrderID string) *AlphaWithdrawService {
	s.clientOrderID = &clientOrderID
	return s
}

// RecvWindow set recvWindow
func (s *AlphaWithdrawService) RecvWindow(recvWindow int64) *AlphaWithdrawService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *AlphaWithdrawService) Do(ctx context.Context, opts ...RequestOption) (*AlphaWithdrawResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/capital/alpha-withdraw/apply",
		secType:  secTypeSigned,
	}
	r.setFormParam("network", s.network)
	r.setFormParam("alphaId", s.alphaID)
	r.setFormParam("contractAddress", s.contractAddress)
	r.setFormParam("address", s.address)
	if s.addressTag != nil {
		r.setFormParam("addressTag", *s.addressTag)
	}
	r.setFormParam("amount", s.amount)
	if s.clientOrderID != nil {
		r.setFormParam("clientOrderId", *s.clientOrderID)
	}
	if s.recvWindow != nil {
		r.setFormParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res AlphaWithdrawResponse
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAlphaWithdrawHistoryService fetch withdraw history
type GetAlphaWithdrawHistoryService struct {
	c             *Client
	alphaID       *string
	clientOrderID *string
	status        *int
	startTime     *int64
	endTime       *int64
	offset        *int
	limit         *int
	idList        *string
	recvWindow    *int64
}

// AlphaID set alphaId
func (s *GetAlphaWithdrawHistoryService) AlphaID(alphaID string) *GetAlphaWithdrawHistoryService {
	s.alphaID = &alphaID
	return s
}

// ClientOrderID set clientOrderId
func (s *GetAlphaWithdrawHistoryService) ClientOrderID(clientOrderID string) *GetAlphaWithdrawHistoryService {
	s.clientOrderID = &clientOrderID
	return s
}

// Status set status
func (s *GetAlphaWithdrawHistoryService) Status(status int) *GetAlphaWithdrawHistoryService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *GetAlphaWithdrawHistoryService) StartTime(startTime int64) *GetAlphaWithdrawHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetAlphaWithdrawHistoryService) EndTime(endTime int64) *GetAlphaWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *GetAlphaWithdrawHistoryService) Offset(offset int) *GetAlphaWithdrawHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *GetAlphaWithdrawHistoryService) Limit(limit int) *GetAlphaWithdrawHistoryService {
	s.limit = &limit
	return s
}

// IDList set idList
func (s *GetAlphaWithdrawHistoryService) IDList(idList string) *GetAlphaWithdrawHistoryService {
	s.idList = &idList
	return s
}

// RecvWindow set recvWindow
func (s *GetAlphaWithdrawHistoryService) RecvWindow(recvWindow int64) *GetAlphaWithdrawHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetAlphaWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AlphaWithdrawHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/alpha-withdraw/history",
		secType:  secTypeSigned,
	}
	if s.alphaID != nil {
		r.setParam("alphaId", *s.alphaID)
	}
	if s.clientOrderID != nil {
		r.setParam("clientOrderId", *s.clientOrderID)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.idList != nil {
		r.setParam("idList", *s.idList)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []AlphaWithdrawHistory
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetAlphaDepositHistoryService fetch deposit history
type GetAlphaDepositHistoryService struct {
	c             *Client
	alphaID       *string
	txID          *string
	status        *int
	startTime     *int64
	endTime       *int64
	includeSource *bool
	offset        *int
	limit         *int
	recvWindow    *int64
}

// AlphaID set alphaId
func (s *GetAlphaDepositHistoryService) AlphaID(alphaID string) *GetAlphaDepositHistoryService {
	s.alphaID = &alphaID
	return s
}

// TxID set txId
func (s *GetAlphaDepositHistoryService) TxID(txID string) *GetAlphaDepositHistoryService {
	s.txID = &txID
	return s
}

// Status set status
func (s *GetAlphaDepositHistoryService) Status(status int) *GetAlphaDepositHistoryService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *GetAlphaDepositHistoryService) StartTime(startTime int64) *GetAlphaDepositHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetAlphaDepositHistoryService) EndTime(endTime int64) *GetAlphaDepositHistoryService {
	s.endTime = &endTime
	return s
}

// IncludeSource set includeSource
func (s *GetAlphaDepositHistoryService) IncludeSource(includeSource bool) *GetAlphaDepositHistoryService {
	s.includeSource = &includeSource
	return s
}

// Offset set offset
func (s *GetAlphaDepositHistoryService) Offset(offset int) *GetAlphaDepositHistoryService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *GetAlphaDepositHistoryService) Limit(limit int) *GetAlphaDepositHistoryService {
	s.limit = &limit
	return s
}

// RecvWindow set recvWindow
func (s *GetAlphaDepositHistoryService) RecvWindow(recvWindow int64) *GetAlphaDepositHistoryService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetAlphaDepositHistoryService) Do(ctx context.Context, opts ...RequestOption) ([]AlphaDepositHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/alpha-deposit/history",
		secType:  secTypeSigned,
	}
	if s.alphaID != nil {
		r.setParam("alphaId", *s.alphaID)
	}
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.includeSource != nil {
		r.setParam("includeSource", *s.includeSource)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
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
	var res []AlphaDepositHistory
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetAlphaDepositAddressService fetch deposit address
type GetAlphaDepositAddressService struct {
	c          *Client
	network    string
	recvWindow *int64
}

// Network set network
func (s *GetAlphaDepositAddressService) Network(network string) *GetAlphaDepositAddressService {
	s.network = network
	return s
}

// RecvWindow set recvWindow
func (s *GetAlphaDepositAddressService) RecvWindow(recvWindow int64) *GetAlphaDepositAddressService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetAlphaDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (*AlphaDepositAddress, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/alpha-deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("network", s.network)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res AlphaDepositAddress
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
