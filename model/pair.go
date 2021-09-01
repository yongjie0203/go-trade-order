package model

// PairGroup 交易对分组信息
type PairGroup struct {
	Id int `json:"id" gorm:"primary_key"`
	//分组名称
	GroupName string `json:"group_name"`
	//分组类别 现货 合约
	GroupType string `json:"group_type"`
	//创建时间
	Time int64 `json:"time"`
	//分组状态 0 无效 1 有效
	Status int `json:"status" gorm:"default:1"`
	//标签
	Tags string `json:"tags"`
	//删除标识 1：删除 0：正常
	IsDelete int `json:"is_delete" gorm:"default:0"`
	//删除时间
	DeleteTime int64 `json:"delete_time"`
}

// Pair 交易对信息
type Pair struct {
	Id
	PariName string `json:"pari_name"`
	GroupId  int64  `json:"group_id"`
	//状态 0 无效 1 有效上架
	Status int `json:"status" gorm:"default:1"`
	//标签
	Tags          string `json:"tags"`
	ValuedTokenId int64  `json:"valued_token_id"`
	//被计价的token name
	ValuedTokenName string `json:"valued_token_name"`

	ValuationTokenId int64 `json:"valuation_token_id"`
	//计价token name
	ValuationTokenName string `json:"valuation_token_name"`

	PriceCurrent float64 `json:"price_current"`

	Tail
}
