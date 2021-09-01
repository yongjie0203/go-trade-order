package reg

import "yingyi.cn/go-trade/trade/handler"

func SetupRouter() {

	group := New()
	group.On("login", handler.Login)

}
