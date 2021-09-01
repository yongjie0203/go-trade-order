package model

// Token token信息
type Token struct {
	Id
	MainNetId int64 `json:"main_net_id"`
	//是否主网token如Eth为以太坊主网 1 主网 0 非主网
	IsMainNet int8 `json:"is_main_net" gorm:"default:0"`
	//1正常
	Status      int    `json:"status"`
	Price       int64  `json:"price"`
	Decimals    uint8  `json:"decimals"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply uint64 `json:"total_supply"`
	//合约地址
	Address string `json:"address"`
	Tail
}
