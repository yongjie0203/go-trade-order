package handler

import (
	"fmt"
	"github.com/yongjie0203/go-trade-order/websockets"
)

func Login(c *websockets.Client, message *websockets.Message) {
	fmt.Println("login request ")
	c.SendMessage([]byte("login success"))
}
