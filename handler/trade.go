package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"yingyi.cn/go-trade-core/core/order"
	"yingyi.cn/go-trade/trade/db"
	"yingyi.cn/go-trade/trade/pair"
	"yingyi.cn/go-trade/trade/request"
)

func Order(ctx *gin.Context) {
	order := new(request.Order)
	rand.Seed(time.Now().UnixNano())
	order.Num = rand.Float64() * 10
	order.Price = rand.Float64() * 100
	ctx.BindJSON(&order)
	item := BuildOrder(order)
	db.Coon.Create(item)
	quevePqir := pair.Pairs["btc"]
	if item.Direction == 1 {
		quevePqir.Buy.Insert(item)
	} else {
		quevePqir.Sell.Insert(item)
	}

	fmt.Printf("sell.size:%d\n ", quevePqir.Sell.Size())
	fmt.Printf("y.size:%d\n ", quevePqir.Buy.Size())
	ctx.JSON(200, item)
}

func BuildOrder(request *request.Order) order.Order {
	var item = order.Order{}
	item.Time = time.Now().UnixNano()
	item.UID = 0
	item.OID = NewOID()
	item.Price = request.Price
	item.Num = request.Num
	item.Direction = randomDirection()
	item.PairId = request.PairId
	item.TokenPair = request.TokenPair
	item.Priority = float64(time.Now().Unix())
	return item
}

func NewOID() int64 {
	return db.NextId("")
}

func randomDirection() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}
