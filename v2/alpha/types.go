package alpha

// TokenConfig represents token configuration info
type TokenConfig struct {
	Network              string `json:"network"`
	Coin                 string `json:"coin"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	EntityTag            string `json:"entityTag"`
	IsDefault            bool   `json:"isDefault"`
	DepositEnable        bool   `json:"depositEnable"`
	WithdrawEnable       bool   `json:"withdrawEnable"`
	DepositDesc          string `json:"depositDesc"`
	WithdrawDesc         string `json:"withdrawDesc"`
	SpecialDepositTips   string `json:"specialDepositTips"`
	SpecialWithdrawTips  string `json:"specialWithdrawTips"`
	AddressRegex         string `json:"addressRegex"`
	AddressRule          string `json:"addressRule"`
	MemoRegex            string `json:"memoRegex"`
	WithdrawFee          string `json:"withdrawFee"`
	WithdrawMin          string `json:"withdrawMin"`
	WithdrawMax          string `json:"withdrawMax"`
	DepositDust          string `json:"depositDust"`
	MinConfirm           int    `json:"minConfirm"`
	UnLockConfirm        int    `json:"unLockConfirm"`
	SameAddress          bool   `json:"sameAddress"`
	EstimatedArrivalTime int    `json:"estimatedArrivalTime"`
	ContractAddressUrl   string `json:"contractAddressUrl"`
}

// ExchangeInfoAsset represents asset in exchange info
type ExchangeInfoAsset struct {
	Asset string `json:"asset"`
}

// ExchangeInfoSymbolFilter represents symbol filter
type ExchangeInfoSymbolFilter struct {
	FilterType        string `json:"filterType"`
	MinPrice          string `json:"minPrice,omitempty"`
	MaxPrice          string `json:"maxPrice,omitempty"`
	TickSize          string `json:"tickSize,omitempty"`
	StepSize          string `json:"stepSize,omitempty"`
	MaxQty            string `json:"maxQty,omitempty"`
	MinQty            string `json:"minQty,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	MinNotional       string `json:"minNotional,omitempty"`
	MaxNotional       string `json:"maxNotional,omitempty"`
	MultiplierDown    string `json:"multiplierDown,omitempty"`
	MultiplierUp      string `json:"multiplierUp,omitempty"`
	BidMultiplierUp   string `json:"bidMultiplierUp,omitempty"`
	AskMultiplierUp   string `json:"askMultiplierUp,omitempty"`
	BidMultiplierDown string `json:"bidMultiplierDown,omitempty"`
	AskMultiplierDown string `json:"askMultiplierDown,omitempty"`
}

// ExchangeInfoSymbol represents symbol in exchange info
type ExchangeInfoSymbol struct {
	Symbol             string                     `json:"symbol"`
	Status             string                     `json:"status"`
	BaseAsset          string                     `json:"baseAsset"`
	QuoteAsset         string                     `json:"quoteAsset"`
	PricePrecision     int                        `json:"pricePrecision"`
	QuantityPrecision  int                        `json:"quantityPrecision"`
	BaseAssetPrecision int                        `json:"baseAssetPrecision"`
	QuotePrecision     int                        `json:"quotePrecision"`
	Filters            []ExchangeInfoSymbolFilter `json:"filters"`
	OrderTypes         []string                   `json:"orderTypes"`
}

// ExchangeInfo represents exchange info response
type ExchangeInfo struct {
	Timezone string               `json:"timezone"`
	Assets   []ExchangeInfoAsset  `json:"assets"`
	Symbols  []ExchangeInfoSymbol `json:"symbols"`
}

// CommissionFee represents commission fee response
type CommissionFee struct {
	BuyerCommission  int `json:"buyerCommission"`
	SellerCommission int `json:"sellerCommission"`
}

// ListenKey represents listen key response
type ListenKey struct {
	ListenKey string `json:"listenKey"`
}

// PlaceOrderResponse represents place order response
type PlaceOrderResponse struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"` // P:Processing, S:Success, F:Failure
}

// CancelOrderResponse represents cancel order response
type CancelOrderResponse struct {
	OrderID     string `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
}

// CancelAllOrdersResponse represents cancel all orders response
type CancelAllOrdersResponse struct {
	Success bool `json:"success"`
}

// Order represents order information
type Order struct {
	OrderID       string `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQuote      string `json:"cumQuote"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	Side          string `json:"side"`
	StopPrice     string `json:"stopPrice"`
	OrigType      string `json:"origType"`
	Time          int64  `json:"time"`
	UpdateTime    int64  `json:"updateTime"`
	OrderListID   string `json:"orderListId"`
	PageID        string `json:"pageId"`
	BaseAsset     string `json:"baseAsset"`
	QuoteAsset    string `json:"quoteAsset"`
}

// Trade represents trade information
type Trade struct {
	Symbol          string `json:"symbol"`
	ID              string `json:"id"`
	OrderID         string `json:"orderId"`
	TradeID         string `json:"tradeId"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	PageID          string `json:"pageId"`
	Buyer           bool   `json:"buyer"`
	BaseAsset       string `json:"baseAsset"`
	QuoteAsset      string `json:"quoteAsset"`
	OrderType       string `json:"orderType"`
	LastTrade       bool   `json:"lastTrade"`
}

// Kline represents kline data
type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           int    `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
	Ignore                   string `json:"ignore"`
}

// Ticker represents 24hr ticker
type Ticker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

// TickerPrice represents ticker price
type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

// AggTrade represents aggregate trade
type AggTrade struct {
	AggTradeID      int64  `json:"a"`
	Price           string `json:"p"`
	Quantity        string `json:"q"`
	FirstTradeID    int64  `json:"f"`
	LastTradeID     int64  `json:"l"`
	IsBuyerMaker    bool   `json:"m"`
	TransactionTime int64  `json:"T"`
}

// BookTicker represents book ticker
type BookTicker struct {
	UpdateID        int64  `json:"u"`
	EventType       string `json:"e"`
	Symbol          string `json:"s"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
	TransactionTime int64  `json:"T"`
	EventTime       int64  `json:"E"`
}

// Depth represents order book depth
type Depth struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Symbol       string     `json:"symbol"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
	T            int64      `json:"t"`
	E            int64      `json:"e"`
}

// AlphaAsset represents alpha wallet asset balance
type AlphaAsset struct {
	ChainID         string `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	AlphaID         string `json:"alphaId"`
	CexAssetCode    string `json:"cexAssetCode"`
	Free            string `json:"free"`
	Freeze          string `json:"freeze"`
	Locked          string `json:"locked"`
	Withdrawing     string `json:"withdrawing"`
	Amount          string `json:"amount"`
	Valuation       string `json:"valuation"`
}

// TokenMapping represents token mapping info
type TokenMapping struct {
	TokenID           string `json:"tokenId"`
	ChainID           string `json:"chainId"`
	ChainIconURL      string `json:"chainIconUrl"`
	ChainName         string `json:"chainName"`
	ContractAddress   string `json:"contractAddress"`
	Name              string `json:"name"`
	Symbol            string `json:"symbol"`
	Price             string `json:"price"`
	PercentChange24h  string `json:"percentChange24h"`
	Volume24h         string `json:"volume24h"`
	MarketCap         string `json:"marketCap"`
	FDV               string `json:"fdv"`
	Liquidity         string `json:"liquidity"`
	TotalSupply       string `json:"totalSupply"`
	CirculatingSupply string `json:"circulatingSupply"`
	Holders           string `json:"holders"`
	Decimals          int    `json:"decimals"`
	ListingCex        bool   `json:"listingCex"`
	HotTag            bool   `json:"hotTag"`
	CexCoinName       string `json:"cexCoinName"`
	CanTransfer       bool   `json:"canTransfer"`
	Denomination      int    `json:"denomination"`
	Offline           bool   `json:"offline"`
	TradeDecimal      int    `json:"tradeDecimal"`
	AlphaID           string `json:"alphaId"`
	Offsell           bool   `json:"offsell"`
	PriceHigh24h      string `json:"priceHigh24h"`
	PriceLow24h       string `json:"priceLow24h"`
	OnlineTge         bool   `json:"onlineTge"`
	OnlineAirdrop     bool   `json:"onlineAirdrop"`
}

// AlphaWithdrawResponse represents alpha withdraw response
type AlphaWithdrawResponse struct {
	ID string `json:"id"`
}

// AlphaWithdrawHistory represents alpha withdraw history
type AlphaWithdrawHistory struct {
	ID              string `json:"id"`
	Network         string `json:"network"`
	AlphaID         string `json:"alphaId"`
	ContractAddress string `json:"contractAddress"`
	CoinName        string `json:"coinName"`
	Address         string `json:"address"`
	AddressTag      string `json:"addressTag"`
	Amount          string `json:"amount"`
	TxID            string `json:"txId"`
	ApplyTime       string `json:"applyTime"`
	CompleteTime    string `json:"completeTime"`
	ConfirmNo       int    `json:"confirmNo"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	Info            string `json:"info"`
}

// AlphaDepositHistory represents alpha deposit history
type AlphaDepositHistory struct {
	ID              string `json:"id"`
	Network         string `json:"network"`
	AlphaID         string `json:"alphaId"`
	ContractAddress string `json:"contractAddress"`
	CoinName        string `json:"coinName"`
	Address         string `json:"address"`
	AddressTag      string `json:"addressTag"`
	Amount          string `json:"amount"`
	TxID            string `json:"txId"`
	CompleteTime    int64  `json:"completeTime"`
	ConfirmationNo  int64  `json:"confirmationNo"`
	InsertTime      int64  `json:"insertTime"`
	SourceAddress   string `json:"sourceAddress"`
	Status          int    `json:"status"`
	UnlockConfirm   int    `json:"unlockConfirm"`
}

// AlphaDepositAddress represents alpha deposit address
type AlphaDepositAddress struct {
	Network string `json:"network"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
}
