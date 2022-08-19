package incoming

import (
	"encoding/binary"
	"errors"
	"hirasawa/bancho/common"
	"hirasawa/bancho/outgoing"
	"io"
	"io/ioutil"
	"log"
)

type HandlerFunc func(p *PacketHeader, ctx *common.Context, r io.Reader) error

type HandlerSet struct {
	Funcs map[PacketID]HandlerFunc
}

func (h *HandlerSet) Handle(ctx *common.Context, r io.Reader) error {
	header := &PacketHeader{}

	err := binary.Read(r, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	handlerFunc, ok := h.Funcs[header.ReadType]
	if !ok {
		return errors.New("Missing handler")
	}

	log.Println("Handling packet", header.ReadType)
	return handlerFunc(header, ctx, r)
}

var MainHandler HandlerSet

func init() {
	MainHandler = HandlerSet{}
	MainHandler.Funcs = map[PacketID]HandlerFunc{}
	MainHandler.Funcs[CHANGE_ACTION] = unimplemented
	MainHandler.Funcs[SEND_PUBLIC_MESSAGE] = unimplemented
	MainHandler.Funcs[LOGOUT] = unimplemented
	MainHandler.Funcs[REQUEST_STATUS_UPDATE] = func(p *PacketHeader, ctx *common.Context, r io.Reader) error {
		ctx.Player.PacketQueue.Write(outgoing.UserStats(ctx.Player.ID, 0, "poop", "poop", 0, 0, 0, 727, 0.69, 123, 234, 345, 72))
		return nil
	}
	MainHandler.Funcs[PING] = doNothing
	MainHandler.Funcs[START_SPECTATING] = unimplemented
	MainHandler.Funcs[STOP_SPECTATING] = unimplemented
	MainHandler.Funcs[SPECTATE_FRAMES] = unimplemented
	MainHandler.Funcs[ERROR_REPORT] = unimplemented
	MainHandler.Funcs[CANT_SPECTATE] = unimplemented
	MainHandler.Funcs[SEND_PRIVATE_MESSAGE] = unimplemented
	MainHandler.Funcs[PART_LOBBY] = unimplemented
	MainHandler.Funcs[JOIN_LOBBY] = unimplemented
	MainHandler.Funcs[CREATE_MATCH] = unimplemented
	MainHandler.Funcs[JOIN_MATCH] = unimplemented
	MainHandler.Funcs[PART_MATCH] = unimplemented
	MainHandler.Funcs[MATCH_CHANGE_SLOT] = unimplemented
	MainHandler.Funcs[MATCH_READY] = unimplemented
	MainHandler.Funcs[MATCH_LOCK] = unimplemented
	MainHandler.Funcs[MATCH_CHANGE_SETTINGS] = unimplemented
	MainHandler.Funcs[MATCH_START] = unimplemented
	MainHandler.Funcs[MATCH_SCORE_UPDATE] = unimplemented
	MainHandler.Funcs[MATCH_COMPLETE] = unimplemented
	MainHandler.Funcs[MATCH_CHANGE_MODS] = unimplemented
	MainHandler.Funcs[MATCH_LOAD_COMPLETE] = unimplemented
	MainHandler.Funcs[MATCH_NO_BEATMAP] = unimplemented
	MainHandler.Funcs[MATCH_NOT_READY] = unimplemented
	MainHandler.Funcs[MATCH_FAILED] = unimplemented
	MainHandler.Funcs[MATCH_HAS_BEATMAP] = unimplemented
	MainHandler.Funcs[MATCH_SKIP_REQUEST] = unimplemented
	MainHandler.Funcs[CHANNEL_JOIN] = unimplemented
	MainHandler.Funcs[BEATMAP_INFO_REQUEST] = unimplemented
	MainHandler.Funcs[MATCH_TRANSFER_HOST] = unimplemented
	MainHandler.Funcs[FRIEND_ADD] = unimplemented
	MainHandler.Funcs[FRIEND_REMOVE] = unimplemented
	MainHandler.Funcs[MATCH_CHANGE_TEAM] = unimplemented
	MainHandler.Funcs[CHANNEL_PART] = unimplemented
	MainHandler.Funcs[RECEIVE_UPDATES] = unimplemented
	MainHandler.Funcs[SET_AWAY_MESSAGE] = unimplemented
	MainHandler.Funcs[IRC_ONLY] = unimplemented
	MainHandler.Funcs[USER_STATS_REQUEST] = func(p *PacketHeader, ctx *common.Context, r io.Reader) error {
		users := readInt32List16(r)
		for u := range users {
			ctx.Player.PacketQueue.Write(outgoing.UserStats(int32(u), 0, "poop", "poop", 0, 0, 0, 727, 0.69, 123, 234, 345, 72))
		}
		return nil
	}
	MainHandler.Funcs[MATCH_INVITE] = unimplemented
	MainHandler.Funcs[MATCH_CHANGE_PASSWORD] = unimplemented
	MainHandler.Funcs[TOURNAMENT_MATCH_INFO_REQUEST] = unimplemented
	MainHandler.Funcs[USER_PRESENCE_REQUEST] = func(p *PacketHeader, ctx *common.Context, r io.Reader) error {
		users := readInt32List16(r)
		for u := range users {
			ctx.Player.PacketQueue.Write(outgoing.UserPresence(int32(u), "Retard", 2, 69, 2, 0, 0, 0, 1))
		}
		return nil
	}
	MainHandler.Funcs[USER_PRESENCE_REQUEST_ALL] = unimplemented
	MainHandler.Funcs[TOGGLE_BLOCK_NON_FRIEND_DMS] = unimplemented
	MainHandler.Funcs[TOURNAMENT_JOIN_MATCH_CHANNEL] = unimplemented
	MainHandler.Funcs[TOURNAMENT_LEAVE_MATCH_CHANNEL] = unimplemented
}

func doNothing(p *PacketHeader, ctx *common.Context, r io.Reader) error {
	_, err := io.CopyN(ioutil.Discard, r, int64(p.Length))
	return err
}

func unimplemented (p *PacketHeader, ctx *common.Context, r io.Reader) error {
	log.Println("Packet handling for", p.ReadType, "is not yet implemented")
	return doNothing(p, ctx, r)
}


