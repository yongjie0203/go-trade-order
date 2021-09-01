package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yongjie0203/go-trade-order/db"
	"github.com/yongjie0203/go-trade-order/handler"
	"github.com/yongjie0203/go-trade-order/pair"
	"github.com/yongjie0203/go-trade-order/reg"
	"github.com/yongjie0203/go-trade-order/websockets"
)

func main() {

	db.InitDB()
	//更新数据结构
	db.DBUpdate()
	pair.InitPairs()
	pair.Start()

	reg.SetupRouter()
	go websockets.SetUpWebSocket()
	websockets.PingTimer()

	ginRouter := gin.Default()

	api := ginRouter.Group("/v1")
	trade := api.Group("/trade")
	{
		trade.GET("/order", handler.Order)
		//trade.POST("/list/:size/:page", controller.GetArticleList)
		//trade.GET("/get/:id", controller.GetArticle)
		//trade.GET("/delete/:id", controller.eteArticle)
	}

	ginRouter.Run()

}
