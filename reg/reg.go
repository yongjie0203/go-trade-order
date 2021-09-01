package reg

import "github.com/yongjie0203/go-trade-order/websockets"

type HandlerGroup struct {
	handlers map[string]websockets.Handler
}

func New() *HandlerGroup {
	engine := new(HandlerGroup)
	engine.handlers = make(map[string]websockets.Handler)
	return engine
}

func (e *HandlerGroup) On(url string, handler websockets.Handler) {
	e.handlers[url] = handler
	websockets.Handlers[url] = handler
}
