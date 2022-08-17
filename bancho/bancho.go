package bancho

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
)

type PacketID uint16

type PacketHeader struct {
	ReadType PacketID
	_        byte // Peppy padding
	Length   uint32
}

func ReadBanchoPacket(r io.Reader) (PacketHeader, error) {
	var header PacketHeader

	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return header, err
	}

	_, err = io.CopyN(ioutil.Discard, r, int64(header.Length))
	if err != nil {
		return header, err
	}

	return header, nil
}

func write(p PacketID, args ...interface{}) []byte {
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
		err = binary.Write(payload, binary.LittleEndian, v)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = binary.Write(buffer, binary.LittleEndian, uint32(payload.Len()))
	if err != nil {
		log.Fatal(err)
	}

	payload.WriteTo(buffer)
	return buffer.Bytes()
}

// Packet 5
func UserID(userID int32) []byte {
	return write(SERVER_USER_ID, userID)
}

// Packet 11
// GUGE: strings might be broken
func UserStats(userID int32, action uint8, infoText string, mapMD5 string, mods int32, mode, uint8, mapID int32,
	rankedScore int64, accuracy float32, plays int32, totalScore int64, globalRank int32, pp int16) []byte {
	return write(SERVER_USER_STATS, userID, action, infoText, mapMD5, mods, mode, mapID, rankedScore, accuracy, plays, totalScore, globalRank, pp)
}

// Packet 12
func Logout(userID int32) []byte {
	return write(SERVER_USER_LOGOUT, userID)
}

// Packet 13
func SpectatorJoined(userID int32) []byte {
	return write(SERVER_SPECTATOR_JOINED, userID)
}

// Packet 14
func SpectatorLeft(userID int32) []byte {
	return write(SERVER_SPECTATOR_LEFT, userID)
}

// Packet 19
func VersionUpdate() []byte {
	return write(SERVER_VERSION_UPDATE)
}

// Packet 22
func SpectatorCantSpectate(userID int32) []byte {
	return write(SERVER_SPECTATOR_CANT_SPECTATE, userID)
}

// Packet 23
func GetAttention() []byte {
	return write(SERVER_GET_ATTENTION)
}

// Packet 24
// GUGE: strings might be broken
func Notification() []byte {
	return write(SERVER_NOTIFICATION)
}

// Packet 28
func DisposeMatch(ID int32) []byte {
	return write(SERVER_DISPOSE_MATCH, ID)
}

// Packet 34
func ToggleBlockNonFriendDM() []byte {
	return write(SERVER_TOGGLE_BLOCK_NON_FRIEND_DMS)
}

// Packet 37
func MatchJoinFail() []byte {
	return write(SERVER_MATCH_JOIN_FAIL)
}

// Packet 42
func FellowSpectatorJoined(userID int32) []byte {
	return write(SERVER_FELLOW_SPECTATOR_JOINED, userID)
}

// Packet 43
func FellowSpectatorLeft(userID int32) []byte {
	return write(SERVER_FELLOW_SPECTATOR_LEFT, userID)
}

// Packet 50
func MatchTransferHost() []byte {
	return write(SERVER_MATCH_TRANSFER_HOST)
}

// Packet 53
func MatchAllPlayersLoaded() []byte {
	return write(SERVER_MATCH_ALL_PLAYERS_LOADED)
}

// Packet 57
func MatchPlayerFailed(slotID int32) []byte {
	return write(SERVER_MATCH_PLAYER_FAILED, slotID)
}

// Packet 58
func MatchComplete() []byte {
	return write(SERVER_MATCH_COMPLETE)
}

// Packet 61
func MatchSkip() []byte {
	return write(SERVER_MATCH_SKIP)
}

// Packet 64
// GUGE: strings might be broken
func ChannelJoin(name string) []byte {
	return write(SERVER_CHANNEL_JOIN_SUCCESS, name)
}

// Packet 66
// GUGE: strings might be broken
func ChannelKick(name string) []byte {
	return write(SERVER_CHANNEL_KICK, name)
}

// Packet 71
func BanchoPrivileges(priv int32) []byte {
	return write(SERVER_PRIVILEGES, priv)
}

// Packet 75
func ProtocolVersion(ver int32) []byte {
	return write(SERVER_PROTOCOL_VERSION, ver)
}

// Packet 81
func MatchPlayerSkipped(userID int32) []byte {
	return write(SERVER_MATCH_PLAYER_SKIPPED, userID)
}

// Packet 83
// GUGE: strings might be broken
// UTCOffset might be broken
// whole thing might be broken
func UserPresence(userID int32, name string, UTCOffset uint8, countryCode uint8, banchoPrivileges uint8, longitude float32,
	latitude float32, globalRank int32) []byte {
	return write(SERVER_USER_PRESENCE, userID, name, UTCOffset, countryCode, banchoPrivileges, longitude, latitude, globalRank)
}

// Packet 86
func RestartServer(ms int32) []byte {
	return write(SERVER_RESTART, ms)
}

// Packet 89
func ChannelInfoEnd() []byte {
	return write(SERVER_CHANNEL_INFO_END)
}

// Packet 91
// GUGE: strings might be broken
func MatchChangePassword(new string) []byte {
	return write(SERVER_MATCH_CHANGE_PASSWORD, new)
}

// Packet 92
func SilenceEnd(delta int32) []byte {
	return write(SERVER_SILENCE_END, delta)
}

// Packet 94
func UserSilenced(userID int32) []byte {
	return write(SERVER_USER_SILENCED, userID)
}

// Packet 95
func UserPresenceSingle(userID int32) []byte {
	return write(SERVER_USER_PRESENCE_SINGLE, userID)
}

// Packet 102
func VersionUpdateForced() []byte {
	return write(SERVER_VERSION_UPDATE_FORCED)
}

// Packet 103
func SwitchServer(t int32) []byte {
	return write(SERVER_SWITCH_SERVER, t)
}

// Packet 104
func AccountRestricted() []byte {
	return write(SERVER_ACCOUNT_RESTRICTED)
}

// Packet 105
// GUGE: strings might be broken
func Rtx(msg string) []byte {
	return write(SERVER_RTX, msg)
}

// Packet 106
func MatchAbort() []byte {
	return write(SERVER_MATCH_ABORT)
}

// Packet 107
// GUGE: strings might be broken
func SwitchTournamentServer(ip string) []byte {
	return write(SERVER_SWITCH_TOURNAMENT_SERVER, ip)
}
