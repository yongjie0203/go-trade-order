package model

// Order 订单信息
type Order struct {
	OID   int64   `json:"oid" gorm:"primary_key;column:oid"`
	UID   int64   `json:"uid" gorm:"index:idx_uid"`
	Price float64 `json:"price"`
	//time added
	Time int64 `json:"time"`
	//buy or sell
	Direction int    `json:"direction"`
	TokenPair string `json:"token_pair"`
	//trade pair Id
	PairId     int     `json:"pair_id" gorm:"index:idx_pair_id"`
	Priority   float64 `json:"priority"`
	Num        float64 `json:"num"`
	Status     int     `json:"status"`
	UpdateTime int64
}

func (o Order) GetPriority() (x float64) {
	return o.Priority
}

// Trade 交易记录
type Trade struct {
	Id
	Oid
	Uid
	Num   float64 `json:"num"`
	Price float64 `json:"price"`
	//完成 取消 1 完成 0 取消
	Status        int    `json:"status"`
	PairName      string `json:"pair_name"`
	ValuedTokenId int64  `json:"valued_token_id"`
	//被计价的TokenSymbol
	ValuedTokenSymbol string `json:"valued_token_symbol"`

	ValuationTokenId int64 `json:"valuation_token_id"`
	//计价TokenSymbol
	ValuationTokenSymbol string `json:"valuation_token_symbol"`
	Tail
}
