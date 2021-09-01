package pair

import (
	"github.com/micro/go-micro/util/log"
	"math"
	"time"
	"yingyi.cn/go-trade-core/core/order"
	"yingyi.cn/go-trade-order/trade/db"
	"yingyi.cn/go-trade-order/trade/model"
)

func OnTransaction(sell, buy order.Order, num float64) {
	log.Log("交易将被处理")
	//保存交易记录
	saveTrade(sell, buy, num)
	//记录最新成交价及成交数量
	saveTransactionPrice(sell, buy, num)

}

//记录最新成交价及成交数量
func saveTransactionPrice(sell, buy order.Order, num float64) {
	transactionPrice := new(model.TransactionPrice)
	transactionPrice.Price = math.Max(sell.Price, buy.Price)
	transactionPrice.Num = num
	transactionPrice.ID = db.NextId("price")
	transactionPrice.PairName = sell.TokenPair
	transactionPrice.PairId = sell.PairId
	transactionPrice.Time = time.Now().UnixNano()
	transactionPrice.BuyOrderId = buy.OID
	transactionPrice.SellOrderId = sell.OID
	db.Coon.Create(transactionPrice)
}

//保存交易记录
func saveTrade(sell, buy order.Order, num float64) {

	sellTrade := buildCompleteTrade(sell, num)
	buyTrade := buildCompleteTrade(buy, num)

	db.Coon.Create(sellTrade)
	db.Coon.Create(buyTrade)
}

func buildCompleteTrade(order order.Order, num float64) *model.Trade {
	trade := new(model.Trade)
	trade.ID = db.NextId("trade")
	trade.Num = num
	trade.UID = order.UID
	trade.OID = order.OID
	trade.Price = order.Price
	trade.Time = time.Now().UnixNano()
	trade.PairName = order.TokenPair
	trade.Status = model.TradeStatusComplete
	return trade
}
