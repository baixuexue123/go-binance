package binance

import "context"

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
		method:   "POST",
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

type ListUniversalTransferService struct {
	c         *Client
	fromEmail *string
	toEmail   *string
	startTime *int64
	endTime   *int64
	page      *int
	limit     *int
}

func (s *ListUniversalTransferService) FromEmail(v string) *ListUniversalTransferService {
	s.fromEmail = &v
	return s
}

func (s *ListUniversalTransferService) ToEmail(v string) *ListUniversalTransferService {
	s.toEmail = &v
	return s
}

func (s *ListUniversalTransferService) StartTime(v int64) *ListUniversalTransferService {
	s.startTime = &v
	return s
}

func (s *ListUniversalTransferService) EndTime(v int64) *ListUniversalTransferService {
	s.endTime = &v
	return s
}

func (s *ListUniversalTransferService) Page(v int) *ListUniversalTransferService {
	s.page = &v
	return s
}

func (s *ListUniversalTransferService) Limit(v int) *ListUniversalTransferService {
	s.limit = &v
	return s
}

func (s *ListUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (res ListUniversalTransferResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/universalTransfer",
		secType:  secTypeSigned,
	}
	if v := s.fromEmail; v != nil {
		r.setParam("fromEmail", *v)
	}
	if v := s.toEmail; v != nil {
		r.setParam("toEmail", *v)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	res.Result = make([]*UniversalTransfer, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

type ListUniversalTransferResponse struct {
	Result []*UniversalTransfer `json:"result"`
}

type UniversalTransfer struct {
	TranId          int64  `json:"tranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	CreateTimeStamp uint64 `json:"createTimeStamp"`
}
