package bancho

import (
	"bytes"
	"encoding/binary"
	"log"
)

type BanchoServerPacketID uint16

const (
	USER_ID                     BanchoServerPacketID = 5
	SEND_MESSAGE                                     = 7
	PONG                                             = 8
	HANDLE_IRC_CHANGE_USERNAME                       = 9
	HANDLE_IRC_QUIT                                  = 10
	USER_STATS                                       = 11
	USER_LOGOUT                                      = 12
	SPECTATOR_JOINED                                 = 13
	SPECTATOR_LEFT                                   = 14
	SPECTATE_FRAMES                                  = 15
	VERSION_UPDATE                                   = 19
	SPECTATOR_CANT_SPECTATE                          = 22
	GET_ATTENTION                                    = 23
	NOTIFICATION                                     = 24
	UPDATE_MATCH                                     = 26
	NEW_MATCH                                        = 27
	DISPOSE_MATCH                                    = 28
	TOGGLE_BLOCK_NON_FRIEND_DMS                      = 34
	MATCH_JOIN_SUCCESS                               = 36
	MATCH_JOIN_FAIL                                  = 37
	FELLOW_SPECTATOR_JOINED                          = 42
	FELLOW_SPECTATOR_LEFT                            = 43
	ALL_PLAYERS_LOADED                               = 45
	MATCH_START                                      = 46
	MATCH_SCORE_UPDATE                               = 48
	MATCH_TRANSFER_HOST                              = 50
	MATCH_ALL_PLAYERS_LOADED                         = 53
	MATCH_PLAYER_FAILED                              = 57
	MATCH_COMPLETE                                   = 58
	MATCH_SKIP                                       = 61
	UNAUTHORIZED                                     = 62
	CHANNEL_JOIN_SUCCESS                             = 64
	CHANNEL_INFO                                     = 65
	CHANNEL_KICK                                     = 66
	CHANNEL_AUTO_JOIN                                = 67
	BEATMAP_INFO_REPLY                               = 69
	PRIVILEGES                                       = 71
	FRIENDS_LIST                                     = 72
	PROTOCOL_VERSION                                 = 75
	MAIN_MENU_ICON                                   = 76
	MONITOR                                          = 80
	MATCH_PLAYER_SKIPPED                             = 81
	USER_PRESENCE                                    = 83
	RESTART                                          = 86
	MATCH_INVITE                                     = 88
	CHANNEL_INFO_END                                 = 89
	MATCH_CHANGE_PASSWORD                            = 91
	SILENCE_END                                      = 92
	USER_SILENCED                                    = 94
	USER_PRESENCE_SINGLE                             = 95
	USER_PRESENCE_BUNDLE                             = 96
	USER_DM_BLOCKED                                  = 100
	TARGET_IS_SILENCED                               = 101
	VERSION_UPDATE_FORCED                            = 102
	SWITCH_SERVER                                    = 103
	ACCOUNT_RESTRICTED                               = 104
	RTX                                              = 105
	MATCH_ABORT                                      = 106
	SWITCH_TOURNAMENT_SERVER                         = 107
)

func write(p BanchoServerPacketID, args ...interface{}) []byte {
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
	return write(USER_ID, userID)
}

// Packet 11
// GUGE: strings might be broken
func UserStats(userID int32, action uint8, infoText string, mapMD5 string, mods int32, mode, uint8, mapID int32,
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
func DisposeMatch(ID int32) []byte {
	return write(DISPOSE_MATCH, ID)
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
// GUGE: strings might be broken
func ChannelJoin(name string) []byte {
	return write(CHANNEL_JOIN_SUCCESS, name)
}

// Packet 66
// GUGE: strings might be broken
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
// GUGE: strings might be broken
// UTCOffset might be broken
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
func MatchChangePassword(new string) []byte {
	return write(MATCH_CHANGE_PASSWORD, new)
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
