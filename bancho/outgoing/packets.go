package outgoing

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Packet 5
func UserID(userID int32) []byte {
	return write(USER_ID, userID)
}

// Packet 8
func Pong() []byte {
	return write(PONG)
}

// Packet 11
func UserStats(userID int32, action uint8, infoText string, mapMD5 string, mods int32, mode uint8, mapID int32,
	rankedScore int64, accuracy float32, plays int32, totalScore int64, globalRank int32, pp int16) []byte {
	return write(USER_STATS, userID, action, infoText, mapMD5, mods, mode, mapID, rankedScore, accuracy, plays, totalScore, globalRank, pp)
}

// Packet 12
func Logout(userID int32) []byte {
	return write(USER_LOGOUT, userID)
}

// Packet 13
func SpectatorJoined(userID int32) []byte {
	return write(SPECTATOR_JOINED, userID)
}

// Packet 14
func SpectatorLeft(userID int32) []byte {
	return write(SPECTATOR_LEFT, userID)
}

// Packet 19
func VersionUpdate() []byte {
	return write(VERSION_UPDATE)
}

// Packet 22
func SpectatorCantSpectate(userID int32) []byte {
	return write(SPECTATOR_CANT_SPECTATE, userID)
}

// Packet 23
func GetAttention() []byte {
	return write(GET_ATTENTION)
}

// Packet 24
// GUGE: strings might be broken
func Notification() []byte {
	return write(NOTIFICATION)
}

// Packet 28
func DisposeMatch(id int32) []byte {
	return write(DISPOSE_MATCH, id)
}

// Packet 34
func ToggleBlockNonFriendDM() []byte {
	return write(TOGGLE_BLOCK_NON_FRIEND_DMS)
}

// Packet 37
func MatchJoinFail() []byte {
	return write(MATCH_JOIN_FAIL)
}

// Packet 42
func FellowSpectatorJoined(userID int32) []byte {
	return write(FELLOW_SPECTATOR_JOINED, userID)
}

// Packet 43
func FellowSpectatorLeft(userID int32) []byte {
	return write(FELLOW_SPECTATOR_LEFT, userID)
}

// Packet 50
func MatchTransferHost() []byte {
	return write(MATCH_TRANSFER_HOST)
}

// Packet 53
func MatchAllPlayersLoaded() []byte {
	return write(MATCH_ALL_PLAYERS_LOADED)
}

// Packet 57
func MatchPlayerFailed(slotID int32) []byte {
	return write(MATCH_PLAYER_FAILED, slotID)
}

// Packet 58
func MatchComplete() []byte {
	return write(MATCH_COMPLETE)
}

// Packet 61
func MatchSkip() []byte {
	return write(MATCH_SKIP)
}

// Packet 64
func ChannelJoin(name string) []byte {
	return write(CHANNEL_JOIN_SUCCESS, name)
}

// Packet 66
func ChannelKick(name string) []byte {
	return write(CHANNEL_KICK, name)
}

// Packet 71
func BanchoPrivileges(priv int32) []byte {
	return write(PRIVILEGES, priv)
}

// Packet 75
func ProtocolVersion(ver int32) []byte {
	return write(PROTOCOL_VERSION, ver)
}

// Packet 81
func MatchPlayerSkipped(userID int32) []byte {
	return write(MATCH_PLAYER_SKIPPED, userID)
}

// Packet 83
// whole thing might be broken
func UserPresence(userID int32, name string, UTCOffset uint8, countryCode uint8, banchoPrivileges uint8, longitude float32,
	latitude float32, globalRank int32) []byte {
	return write(USER_PRESENCE, userID, name, UTCOffset, countryCode, banchoPrivileges, longitude, latitude, globalRank)
}

// Packet 86
func RestartServer(ms int32) []byte {
	return write(RESTART, ms)
}

// Packet 89
func ChannelInfoEnd() []byte {
	return write(CHANNEL_INFO_END)
}

// Packet 91
// GUGE: strings might be broken
func MatchChangePassword(passwd string) []byte {
	return write(MATCH_CHANGE_PASSWORD, passwd)
}

// Packet 92
func SilenceEnd(delta int32) []byte {
	return write(SILENCE_END, delta)
}

// Packet 94
func UserSilenced(userID int32) []byte {
	return write(USER_SILENCED, userID)
}

// Packet 95
func UserPresenceSingle(userID int32) []byte {
	return write(USER_PRESENCE_SINGLE, userID)
}

// Packet 102
func VersionUpdateForced() []byte {
	return write(VERSION_UPDATE_FORCED)
}

// Packet 103
func SwitchServer(t int32) []byte {
	return write(SWITCH_SERVER, t)
}

// Packet 104
func AccountRestricted() []byte {
	return write(ACCOUNT_RESTRICTED)
}

// Packet 105
// GUGE: strings might be broken
func Rtx(msg string) []byte {
	return write(RTX, msg)
}

// Packet 106
func MatchAbort() []byte {
	return write(MATCH_ABORT)
}

// Packet 107
// GUGE: strings might be broken
func SwitchTournamentServer(ip string) []byte {
	return write(SWITCH_TOURNAMENT_SERVER, ip)
}

func write(p OutgoingPacketID, args ...interface{}) []byte {
	buffer := &bytes.Buffer{}

	err := binary.Write(buffer, binary.LittleEndian, p) // 2 Bytes
	if err != nil {
		log.Fatal(err)
	}

	err = binary.Write(buffer, binary.LittleEndian, byte(0)) // 1 Byte Padding???? Peppy?
	if err != nil {
		log.Fatal(err)
	}

	payload := &bytes.Buffer{}
	for _, v := range args {
		switch t := v.(type) {
		case string:
			payload.Write(writeString(t))
		default:
			err = binary.Write(payload, binary.LittleEndian, v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = binary.Write(buffer, binary.LittleEndian, uint32(payload.Len()))
	if err != nil {
		log.Fatal(err)
	}

	payload.WriteTo(buffer)
	return buffer.Bytes()
}

func writeString(str string) []byte {
	if len(str) != 0 {
		b := bytes.Buffer{}
		b.Write([]byte{0x0b})
		str_bytes := []byte(str)
		b.Write(writeUleb128(len(str_bytes)))
		b.Write(str_bytes)
		return b.Bytes()
	} else {
		return []byte{0}
	}
}

func writeUleb128(num int) []byte {
	if num == 0 {
		return []byte{0}
	}

	ret := &bytes.Buffer{}

	for num != 0 {
		binary.Write(ret, binary.LittleEndian, num&0x7F)
		num >>= 7
		if num != 0 {
			ret.Bytes()[ret.Len()-1] |= 0x80
		}
	}

	return ret.Bytes()
}
