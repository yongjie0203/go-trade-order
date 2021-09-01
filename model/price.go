package model

// TransactionPrice 记录历史成交价
type TransactionPrice struct {
	Id
	Time int64 `json:"time"`
	//交易对ID
	PairId int `json:"pair_id"`
	//交易对名称
	PairName string `json:"pair_name"`
	//成交数量
	Num float64 `json:"num"`
	//成交价
	Price       float64 `json:"price"`
	BuyOrderId  int64   `json:"buy_order_id"`
	SellOrderId int64   `json:"sell_order_id"`
}
