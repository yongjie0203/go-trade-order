package reg

import "github.com/yongjie0203/go-trade-order/handler"

func SetupRouter() {

	group := New()
	group.On("login", handler.Login)

}
