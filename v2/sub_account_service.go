package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetSubAccountListService struct {
	c *Client
}

func (s *GetSubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountList, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountList)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountList struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email      string `json:"email"`
	IsFreeze   bool   `json:"isFreeze"`
	UpdateTime uint64 `json:"createTime"`
}

type GetSubAccountAssetsService struct {
	c     *Client
	email *string
}

func (s *GetSubAccountAssetsService) Email(v string) *GetSubAccountAssetsService {
	s.email = &v
	return s
}

func (s *GetSubAccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountAssets, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v3/sub-account/assets",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountAssets)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountAssets struct {
	Balances []Balance `json:"balances"`
}

// CreateUniversalTransferService Universal Transfer (For Master Account)
type CreateUniversalTransferService struct {
	c               *Client
	fromEmail       *string
	toEmail         *string
	fromAccountType *string // "SPOT","USDT_FUTURE","COIN_FUTURE"
	toAccountType   *string // "SPOT","USDT_FUTURE","COIN_FUTURE"
	asset           string
	amount          float64
}

func (s *CreateUniversalTransferService) FromEmail(v string) *CreateUniversalTransferService {
	s.fromEmail = &v
	return s
}

func (s *CreateUniversalTransferService) ToEmail(v string) *CreateUniversalTransferService {
	s.toEmail = &v
	return s
}

func (s *CreateUniversalTransferService) FromAccountType(v string) *CreateUniversalTransferService {
	s.fromAccountType = &v
	return s
}

func (s *CreateUniversalTransferService) ToAccountType(v string) *CreateUniversalTransferService {
	s.toAccountType = &v
	return s
}

func (s *CreateUniversalTransferService) Asset(v string) *CreateUniversalTransferService {
	s.asset = v
	return s
}

func (s *CreateUniversalTransferService) Amount(v float64) *CreateUniversalTransferService {
	s.amount = v
	return s
}

func (s *CreateUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (*CreateUniversalTransferResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if v := s.fromEmail; v != nil {
		r.setParam("fromEmail", *v)
	}
	if v := s.toEmail; v != nil {
		r.setParam("toEmail", *v)
	}
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if v := s.fromAccountType; v != nil {
		r.setParam("fromAccountType", *v)
	}
	if v := s.toAccountType; v != nil {
		r.setParam("toAccountType", *v)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &CreateUniversalTransferResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

type CreateUniversalTransferResponse struct {
	ID int64 `json:"tranId"`
}
