package incoming

import (
	"hirasawa/bancho/common"
	"hirasawa/bancho/outgoing"
)

// Packet 3
type RequestStatusUpdate struct {
	PacketHeader
}

func (RequestStatusUpdate) Handle(ctx *common.Context) {
	ctx.Player.PacketQueue.Write(outgoing.UserStats(69, 0, "Hai", "123", 0, 0, 0, 727, 0.69, 123, 234, 345, 72))
}

// Packet 4
type Ping struct {
	PacketHeader
}

func (Ping) Handle(ctx *common.Context) {
	return
}

// Packet 85
type UserStatsRequest struct {
	PacketHeader
	UserIDs []int32
}

func (UserStatsRequest) Handle(ctx *common.Context) {

}
