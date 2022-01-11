package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adshao/go-binance/v2/common"
)

type GetOpenInterestService struct {
	c      *Client
	symbol *string
}

func (s *GetOpenInterestService) Symbol(symbol string) *GetOpenInterestService {
	s.symbol = &symbol
	return s
}

func (s *GetOpenInterestService) Do(ctx context.Context, opts ...RequestOption) (res *OpenInterest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openInterest",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return res, err
	}
	res = new(OpenInterest)
	err = json.Unmarshal(data, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

type OpenInterest struct {
	Symbol       string `json:"symbol"`
	OpenInterest string `json:"openInterest"`
	Time         uint64 `json:"time"`
}
