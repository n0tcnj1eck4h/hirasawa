package incoming

import (
	"hirasawa/bancho/common"
	"hirasawa/bancho/outgoing"
)

func (Ping) Handle(ctx *common.Context) {
	// Ping doesn't work the way I thought...
	// ctx.PacketQueue.Write(outgoing.Pong())
}

func (RequestStatusUpdate) Handle(ctx *common.Context) {
	ctx.PacketQueue.Write(outgoing.UserStats(69, 0, "Hai", "123", 0, 0, 0, 727, 0.69, 123, 234, 345, 72))
}
