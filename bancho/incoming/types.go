package incoming

import (
	"fmt"
)

type PacketID uint16

type PacketHeader struct {
	ReadType PacketID
	_        byte // Peppy padding
	Length   uint32
}

const (
	CHANGE_ACTION                  PacketID = 0
	SEND_PUBLIC_MESSAGE                     = 1
	LOGOUT                                  = 2
	REQUEST_STATUS_UPDATE                   = 3
	PING                                    = 4
	START_SPECTATING                        = 16
	STOP_SPECTATING                         = 17
	SPECTATE_FRAMES                         = 18
	ERROR_REPORT                            = 20
	CANT_SPECTATE                           = 21
	SEND_PRIVATE_MESSAGE                    = 25
	PART_LOBBY                              = 29
	JOIN_LOBBY                              = 30
	CREATE_MATCH                            = 31
	JOIN_MATCH                              = 32
	PART_MATCH                              = 33
	MATCH_CHANGE_SLOT                       = 38
	MATCH_READY                             = 39
	MATCH_LOCK                              = 40
	MATCH_CHANGE_SETTINGS                   = 41
	MATCH_START                             = 44
	MATCH_SCORE_UPDATE                      = 47
	MATCH_COMPLETE                          = 49
	MATCH_CHANGE_MODS                       = 51
	MATCH_LOAD_COMPLETE                     = 52
	MATCH_NO_BEATMAP                        = 54
	MATCH_NOT_READY                         = 55
	MATCH_FAILED                            = 56
	MATCH_HAS_BEATMAP                       = 59
	MATCH_SKIP_REQUEST                      = 60
	CHANNEL_JOIN                            = 63
	BEATMAP_INFO_REQUEST                    = 68
	MATCH_TRANSFER_HOST                     = 70
	FRIEND_ADD                              = 73
	FRIEND_REMOVE                           = 74
	MATCH_CHANGE_TEAM                       = 77
	CHANNEL_PART                            = 78
	RECEIVE_UPDATES                         = 79
	SET_AWAY_MESSAGE                        = 82
	IRC_ONLY                                = 84
	USER_STATS_REQUEST                      = 85
	MATCH_INVITE                            = 87
	MATCH_CHANGE_PASSWORD                   = 90
	TOURNAMENT_MATCH_INFO_REQUEST           = 93
	USER_PRESENCE_REQUEST                   = 97
	USER_PRESENCE_REQUEST_ALL               = 98
	TOGGLE_BLOCK_NON_FRIEND_DMS             = 99
	TOURNAMENT_JOIN_MATCH_CHANNEL           = 108
	TOURNAMENT_LEAVE_MATCH_CHANNEL          = 109
)

func (p PacketID) String() string {
	switch p {
	case CHANGE_ACTION:
		return "CHANGE_ACTION"
	case SEND_PUBLIC_MESSAGE:
		return "SEND_PUBLIC_MESSAGE"
	case LOGOUT:
		return "LOGOUT"
	case REQUEST_STATUS_UPDATE:
		return "REQUEST_STATUS_UPDATE"
	case PING:
		return "PING"
	case START_SPECTATING:
		return "START_SPECTATING"
	case STOP_SPECTATING:
		return "STOP_SPECTATING"
	case SPECTATE_FRAMES:
		return "SPECTATE_FRAMES"
	case ERROR_REPORT:
		return "ERROR_REPORT"
	case CANT_SPECTATE:
		return "CANT_SPECTATE"
	case SEND_PRIVATE_MESSAGE:
		return "SEND_PRIVATE_MESSAGE"
	case PART_LOBBY:
		return "PART_LOBBY"
	case JOIN_LOBBY:
		return "JOIN_LOBBY"
	case CREATE_MATCH:
		return "CREATE_MATCH"
	case JOIN_MATCH:
		return "JOIN_MATCH"
	case PART_MATCH:
		return "PART_MATCH"
	case MATCH_CHANGE_SLOT:
		return "MATCH_CHANGE_SLOT"
	case MATCH_READY:
		return "MATCH_READY"
	case MATCH_LOCK:
		return "MATCH_LOCK"
	case MATCH_CHANGE_SETTINGS:
		return "MATCH_CHANGE_SETTINGS"
	case MATCH_START:
		return "MATCH_START"
	case MATCH_SCORE_UPDATE:
		return "MATCH_SCORE_UPDATE"
	case MATCH_COMPLETE:
		return "MATCH_COMPLETE"
	case MATCH_CHANGE_MODS:
		return "MATCH_CHANGE_MODS"
	case MATCH_LOAD_COMPLETE:
		return "MATCH_LOAD_COMPLETE"
	case MATCH_NO_BEATMAP:
		return "MATCH_NO_BEATMAP"
	case MATCH_NOT_READY:
		return "MATCH_NOT_READY"
	case MATCH_FAILED:
		return "MATCH_FAILED"
	case MATCH_HAS_BEATMAP:
		return "MATCH_HAS_BEATMAP"
	case MATCH_SKIP_REQUEST:
		return "MATCH_SKIP_REQUEST"
	case CHANNEL_JOIN:
		return "CHANNEL_JOIN"
	case BEATMAP_INFO_REQUEST:
		return "BEATMAP_INFO_REQUEST"
	case MATCH_TRANSFER_HOST:
		return "MATCH_TRANSFER_HOST"
	case FRIEND_ADD:
		return "FRIEND_ADD"
	case FRIEND_REMOVE:
		return "FRIEND_REMOVE"
	case MATCH_CHANGE_TEAM:
		return "MATCH_CHANGE_TEAM"
	case CHANNEL_PART:
		return "CHANNEL_PART"
	case RECEIVE_UPDATES:
		return "RECEIVE_UPDATES"
	case SET_AWAY_MESSAGE:
		return "SET_AWAY_MESSAGE"
	case IRC_ONLY:
		return "IRC_ONLY"
	case USER_STATS_REQUEST:
		return "USER_STATS_REQUEST"
	case MATCH_INVITE:
		return "MATCH_INVITE"
	case MATCH_CHANGE_PASSWORD:
		return "MATCH_CHANGE_PASSWORD"
	case TOURNAMENT_MATCH_INFO_REQUEST:
		return "TOURNAMENT_MATCH_INFO_REQUEST"
	case USER_PRESENCE_REQUEST:
		return "USER_PRESENCE_REQUEST"
	case USER_PRESENCE_REQUEST_ALL:
		return "USER_PRESENCE_REQUEST_ALL"
	case TOGGLE_BLOCK_NON_FRIEND_DMS:
		return "TOGGLE_BLOCK_NON_FRIEND_DMS"
	case TOURNAMENT_JOIN_MATCH_CHANNEL:
		return "TOURNAMENT_JOIN_MATCH_CHANNEL"
	case TOURNAMENT_LEAVE_MATCH_CHANNEL:
		return "TOURNAMENT_LEAVE_MATCH_CHANNEL"
	default:
		return fmt.Sprintf("UNKNOWN (%d)", uint16(p))
	}
}
