package server

import (
	"encoding/binary"
	"errors"
	"hirasawa/bancho/outgoing"
	"hirasawa/bancho/userstore"
	"io"
	"io/ioutil"
	"log"
)

type context struct {
	Player *userstore.Player
	Packet *PacketHeader
	Server *BanchoServer
}

type HandlerFunc func(ctx *context, r io.Reader) error

type Handlers struct {
	Funcs map[PacketID]HandlerFunc
}

func NewHandlers() (h *Handlers) {
	h = &Handlers{}
	h.Funcs = make(map[PacketID]HandlerFunc)
	return h
}

func (h *Handlers) Dispatch(ctx *context, r io.Reader) error {
	header := &PacketHeader{}

	err := binary.Read(r, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	ctx.Packet = header // this is disgusting and confusing

	handlerFunc, ok := h.Funcs[header.ReadType]
	if !ok {
		skip(ctx, r)
		return errors.New("Missing handler")
	}

	log.Println("Handling packet", header.ReadType)
	return handlerFunc(ctx, r)
}

func (h Handlers) InitDefaultHandlers() {
	h.Funcs[CHANGE_ACTION] = func(ctx *context, r io.Reader) error {
		var action uint8
		err := binary.Read(r, binary.LittleEndian, &action)
		if err != nil {
			return err
		}

		infoText, err := readString(r)
		if err != nil {
			return err
		}

		mapHash, err := readString(r)
		if err != nil {
			return err
		}

		var mods uint32
		err = binary.Read(r, binary.LittleEndian, &mods)
		if err != nil {
			return err
		}

		var mode uint8
		err = binary.Read(r, binary.LittleEndian, &mode)
		if err != nil {
			return err
		}

		var mapID int32
		err = binary.Read(r, binary.LittleEndian, &mapID)
		if err != nil {
			return err
		}

		ctx.Player.Session.Status.Action = action
		ctx.Player.Session.Status.InfoText = infoText
		ctx.Player.Session.Status.MapHash = mapHash
		ctx.Player.Session.Status.Mods = mods
		ctx.Player.Session.Status.Mode = mode
		ctx.Player.Session.Status.MapID = mapID

		return nil
	}

	h.Funcs[SEND_PUBLIC_MESSAGE] = unimplemented
	h.Funcs[LOGOUT] = func(ctx *context, r io.Reader) error {
		ctx.Server.PlayerStore.Remove(ctx.Player)
		skip(ctx, r)
		return nil
	}

	h.Funcs[REQUEST_STATUS_UPDATE] = func(ctx *context, r io.Reader) error {
		p := ctx.Player
		p.Session.PacketQueue.Write(outgoing.UserStatsPlayer(p))
		return nil
	}

	h.Funcs[PING] = skip
	h.Funcs[START_SPECTATING] = unimplemented
	h.Funcs[STOP_SPECTATING] = unimplemented
	h.Funcs[SPECTATE_FRAMES] = unimplemented
	h.Funcs[ERROR_REPORT] = unimplemented
	h.Funcs[CANT_SPECTATE] = unimplemented
	h.Funcs[SEND_PRIVATE_MESSAGE] = unimplemented
	h.Funcs[PART_LOBBY] = unimplemented
	h.Funcs[JOIN_LOBBY] = unimplemented
	h.Funcs[CREATE_MATCH] = unimplemented
	h.Funcs[JOIN_MATCH] = unimplemented
	h.Funcs[PART_MATCH] = unimplemented
	h.Funcs[MATCH_CHANGE_SLOT] = unimplemented
	h.Funcs[MATCH_READY] = unimplemented
	h.Funcs[MATCH_LOCK] = unimplemented
	h.Funcs[MATCH_CHANGE_SETTINGS] = unimplemented
	h.Funcs[MATCH_START] = unimplemented
	h.Funcs[MATCH_SCORE_UPDATE] = unimplemented
	h.Funcs[MATCH_COMPLETE] = unimplemented
	h.Funcs[MATCH_CHANGE_MODS] = unimplemented
	h.Funcs[MATCH_LOAD_COMPLETE] = unimplemented
	h.Funcs[MATCH_NO_BEATMAP] = unimplemented
	h.Funcs[MATCH_NOT_READY] = unimplemented
	h.Funcs[MATCH_FAILED] = unimplemented
	h.Funcs[MATCH_HAS_BEATMAP] = unimplemented
	h.Funcs[MATCH_SKIP_REQUEST] = unimplemented
	h.Funcs[CHANNEL_JOIN] = unimplemented
	h.Funcs[BEATMAP_INFO_REQUEST] = unimplemented
	h.Funcs[MATCH_TRANSFER_HOST] = unimplemented
	h.Funcs[FRIEND_ADD] = func(ctx *context, r io.Reader) error {
		var userId int32
		err := binary.Read(r, binary.LittleEndian, &userId)
		if err != nil {
			return err
		}

		// TODO: ADD FRIEND
		return nil
	}

	h.Funcs[FRIEND_REMOVE] = unimplemented
	h.Funcs[MATCH_CHANGE_TEAM] = unimplemented
	h.Funcs[CHANNEL_PART] = unimplemented
	h.Funcs[RECEIVE_UPDATES] = unimplemented
	h.Funcs[SET_AWAY_MESSAGE] = unimplemented
	h.Funcs[IRC_ONLY] = unimplemented
	h.Funcs[USER_STATS_REQUEST] = func(ctx *context, r io.Reader) error {
		users, err := readInt32List16(r)
		if err != nil {
			return err
		}

		for u := range users {
			player, ok := ctx.Server.PlayerStore.FromID(int32(u))
			if ok && player.Online() {
				player.Session.PacketQueue.Write(outgoing.UserStatsPlayer(player))
			}
		}

		return nil
	}

	h.Funcs[MATCH_INVITE] = unimplemented
	h.Funcs[MATCH_CHANGE_PASSWORD] = unimplemented
	h.Funcs[TOURNAMENT_MATCH_INFO_REQUEST] = unimplemented
	h.Funcs[USER_PRESENCE_REQUEST] = func(ctx *context, r io.Reader) error {
		users, err := readInt32List16(r)
		if err != nil {
			return err
		}

		for u := range users {
			player, ok := ctx.Server.PlayerStore.FromID(int32(u))
			if ok && player.Online() {
				player.Session.PacketQueue.Write(outgoing.UserPresencePlayer(player))
			}
		}

		return nil
	}

	h.Funcs[USER_PRESENCE_REQUEST_ALL] = unimplemented
	h.Funcs[TOGGLE_BLOCK_NON_FRIEND_DMS] = unimplemented
	h.Funcs[TOURNAMENT_JOIN_MATCH_CHANNEL] = unimplemented
	h.Funcs[TOURNAMENT_LEAVE_MATCH_CHANNEL] = unimplemented
}

func skip(ctx *context, r io.Reader) error {
	_, err := io.CopyN(ioutil.Discard, r, int64(ctx.Packet.Length))
	return err
}

func unimplemented(ctx *context, r io.Reader) error {
	log.Println("Packet handling for", ctx.Packet.ReadType, "is not yet implemented")
	return skip(ctx, r)
}
