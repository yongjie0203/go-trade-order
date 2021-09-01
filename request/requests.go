package request

type Order struct {
	OID   int64   `json:"oid"`
	UID   int64   `json:"uid"`
	Price float64 `json:"price"`
	//time added
	Time int64 `json:"time"`
	//buy or sell
	Direction int    `json:"direction"`
	TokenPair string `json:"token_pair"`
	//trade pair Id
	PairId   int     `json:"pair_id"`
	Priority float64 `json:"priority"`
	Num      float64 `json:"num"`
}

func (o Order) GetPriority() (x float64) {
	return o.Priority
}
