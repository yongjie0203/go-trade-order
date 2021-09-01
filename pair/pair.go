package pair

import (
	"fmt"
	"github.com/yongjie0203/go-trade-core/queue"
	"sync"
)

type QueuePair struct {
	Buy  *queue.PriorityQueue
	Sell *queue.PriorityQueue
}

var Pairs = make(map[string]QueuePair)

func InitPairs() {
	var buy = new(queue.PriorityQueue)
	var sell = new(queue.PriorityQueue)
	buy.Name("buy").SetLock(new(sync.RWMutex))
	buy.SetHeadLock(new(sync.RWMutex))
	buy.SetTailLock(new(sync.RWMutex))

	sell.Name("ell").SetLock(new(sync.RWMutex))
	sell.SetHeadLock(new(sync.RWMutex))
	sell.SetTailLock(new(sync.RWMutex))
	pair := QueuePair{Buy: buy, Sell: sell}
	Pairs["btc"] = pair
}

func RegisterTransactionHandler() {
	queue.RegistTransactionHandler("handler11", OnTransaction)
}

func Start() {
	//注册交易处理器
	RegisterTransactionHandler()
	for key, pair := range Pairs {
		go queue.Transaction(pair.Sell, pair.Buy)
		fmt.Printf("start pair: %s", key)
	}
}
