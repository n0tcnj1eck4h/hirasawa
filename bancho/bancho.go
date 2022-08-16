package bancho

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
)

type PacketID uint16

const (
	SERVER_USER_ID                     PacketID = 5
	SERVER_SEND_MESSAGE                         = 7
	SERVER_PONG                                 = 8
	SERVER_HANDLE_IRC_CHANGE_USERNAME           = 9
	SERVER_HANDLE_IRC_QUIT                      = 10
	SERVER_USER_STATS                           = 11
	SERVER_USER_LOGOUT                          = 12
	SERVER_SPECTATOR_JOINED                     = 13
	SERVER_SPECTATOR_LEFT                       = 14
	SERVER_SPECTATE_FRAMES                      = 15
	SERVER_VERSION_UPDATE                       = 19
	SERVER_SPECTATOR_CANT_SPECTATE              = 22
	SERVER_GET_ATTENTION                        = 23
	SERVER_NOTIFICATION                         = 24
	SERVER_UPDATE_MATCH                         = 26
	SERVER_NEW_MATCH                            = 27
	SERVER_DISPOSE_MATCH                        = 28
	SERVER_TOGGLE_BLOCK_NON_FRIEND_DMS          = 34
	SERVER_MATCH_JOIN_SUCCESS                   = 36
	SERVER_MATCH_JOIN_FAIL                      = 37
	SERVER_FELLOW_SPECTATOR_JOINED              = 42
	SERVER_FELLOW_SPECTATOR_LEFT                = 43
	SERVER_ALL_PLAYERS_LOADED                   = 45
	SERVER_MATCH_START                          = 46
	SERVER_MATCH_SCORE_UPDATE                   = 48
	SERVER_MATCH_TRANSFER_HOST                  = 50
	SERVER_MATCH_ALL_PLAYERS_LOADED             = 53
	SERVER_MATCH_PLAYER_FAILED                  = 57
	SERVER_MATCH_COMPLETE                       = 58
	SERVER_MATCH_SKIP                           = 61
	SERVER_UNAUTHORIZED                         = 62
	SERVER_CHANNEL_JOIN_SUCCESS                 = 64
	SERVER_CHANNEL_INFO                         = 65
	SERVER_CHANNEL_KICK                         = 66
	SERVER_CHANNEL_AUTO_JOIN                    = 67
	SERVER_BEATMAP_INFO_REPLY                   = 69
	SERVER_PRIVILEGES                           = 71
	SERVER_FRIENDS_LIST                         = 72
	SERVER_PROTOCOL_VERSION                     = 75
	SERVER_MAIN_MENU_ICON                       = 76
	SERVER_MONITOR                              = 80
	SERVER_MATCH_PLAYER_SKIPPED                 = 81
	SERVER_USER_PRESENCE                        = 83
	SERVER_RESTART                              = 86
	SERVER_MATCH_INVITE                         = 88
	SERVER_CHANNEL_INFO_END                     = 89
	SERVER_MATCH_CHANGE_PASSWORD                = 91
	SERVER_SILENCE_END                          = 92
	SERVER_USER_SILENCED                        = 94
	SERVER_USER_PRESENCE_SINGLE                 = 95
	SERVER_USER_PRESENCE_BUNDLE                 = 96
	SERVER_USER_DM_BLOCKED                      = 100
	SERVER_TARGET_IS_SILENCED                   = 101
	SERVER_VERSION_UPDATE_FORCED                = 102
	SERVER_SWITCH_SERVER                        = 103
	SERVER_ACCOUNT_RESTRICTED                   = 104
	SERVER_RTX                                  = 105
	SERVER_MATCH_ABORT                          = 106
	SERVER_SWITCH_TOURNAMENT_SERVER             = 107

	CLIENT_CHANGE_ACTION                  = 0
	CLIENT_SEND_PUBLIC_MESSAGE            = 1
	CLIENT_LOGOUT                         = 2
	CLIENT_REQUEST_STATUS_UPDATE          = 3
	CLIENT_PING                           = 4
	CLIENT_START_SPECTATING               = 16
	CLIENT_STOP_SPECTATING                = 17
	CLIENT_SPECTATE_FRAMES                = 18
	CLIENT_ERROR_REPORT                   = 20
	CLIENT_CANT_SPECTATE                  = 21
	CLIENT_SEND_PRIVATE_MESSAGE           = 25
	CLIENT_PART_LOBBY                     = 29
	CLIENT_JOIN_LOBBY                     = 30
	CLIENT_CREATE_MATCH                   = 31
	CLIENT_JOIN_MATCH                     = 32
	CLIENT_PART_MATCH                     = 33
	CLIENT_MATCH_CHANGE_SLOT              = 38
	CLIENT_MATCH_READY                    = 39
	CLIENT_MATCH_LOCK                     = 40
	CLIENT_MATCH_CHANGE_SETTINGS          = 41
	CLIENT_MATCH_START                    = 44
	CLIENT_MATCH_SCORE_UPDATE             = 47
	CLIENT_MATCH_COMPLETE                 = 49
	CLIENT_MATCH_CHANGE_MODS              = 51
	CLIENT_MATCH_LOAD_COMPLETE            = 52
	CLIENT_MATCH_NO_BEATMAP               = 54
	CLIENT_MATCH_NOT_READY                = 55
	CLIENT_MATCH_FAILED                   = 56
	CLIENT_MATCH_HAS_BEATMAP              = 59
	CLIENT_MATCH_SKIP_REQUEST             = 60
	CLIENT_CHANNEL_JOIN                   = 63
	CLIENT_BEATMAP_INFO_REQUEST           = 68
	CLIENT_MATCH_TRANSFER_HOST            = 70
	CLIENT_FRIEND_ADD                     = 73
	CLIENT_FRIEND_REMOVE                  = 74
	CLIENT_MATCH_CHANGE_TEAM              = 77
	CLIENT_CHANNEL_PART                   = 78
	CLIENT_RECEIVE_UPDATES                = 79
	CLIENT_SET_AWAY_MESSAGE               = 82
	CLIENT_IRC_ONLY                       = 84
	CLIENT_USER_STATS_REQUEST             = 85
	CLIENT_MATCH_INVITE                   = 87
	CLIENT_MATCH_CHANGE_PASSWORD          = 90
	CLIENT_TOURNAMENT_MATCH_INFO_REQUEST  = 93
	CLIENT_USER_PRESENCE_REQUEST          = 97
	CLIENT_USER_PRESENCE_REQUEST_ALL      = 98
	CLIENT_TOGGLE_BLOCK_NON_FRIEND_DMS    = 99
	CLIENT_TOURNAMENT_JOIN_MATCH_CHANNEL  = 108
	CLIENT_TOURNAMENT_LEAVE_MATCH_CHANNEL = 109
)

type PacketHeader struct {
	ReadType PacketID
	_        byte // Peppy padding
	Length   uint32
}

func ReadBanchoPacket(r io.Reader) (PacketHeader, error) {
	var header PacketHeader

	log.Println("Parsing packet header...")

	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return header, err
	}

	log.Printf("Read packet header: %#v\n", header)
	log.Println("Skipping ", header.Length, " bytes...")

	io.CopyN(ioutil.Discard, r, int64(header.Length))

	if err != nil {
		return header, err
	}

	return header, nil
}

func write(p PacketID, args ...interface{}) []byte {
	buffer := bytes.NewBuffer(make([]byte, 0))

	err := binary.Write(buffer, binary.LittleEndian, p) // 2 Bytes
	if err != nil {
		log.Fatal(err)
	}

	err = binary.Write(buffer, binary.LittleEndian, byte(0)) // 1 Byte Padding???? Peppy?
	if err != nil {
		log.Fatal(err)
	}

	payload := bytes.NewBuffer(make([]byte, 0))
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
