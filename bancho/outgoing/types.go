package outgoing

import "fmt"

type OutgoingPacketID uint16

const (
	USER_ID                     OutgoingPacketID = 5
	SEND_MESSAGE                                 = 7
	PONG                                         = 8
	HANDLE_IRC_CHANGE_USERNAME                   = 9
	HANDLE_IRC_QUIT                              = 10
	USER_STATS                                   = 11
	USER_LOGOUT                                  = 12
	SPECTATOR_JOINED                             = 13
	SPECTATOR_LEFT                               = 14
	SPECTATE_FRAMES                              = 15
	VERSION_UPDATE                               = 19
	SPECTATOR_CANT_SPECTATE                      = 22
	GET_ATTENTION                                = 23
	NOTIFICATION                                 = 24
	UPDATE_MATCH                                 = 26
	NEW_MATCH                                    = 27
	DISPOSE_MATCH                                = 28
	TOGGLE_BLOCK_NON_FRIEND_DMS                  = 34
	MATCH_JOIN_SUCCESS                           = 36
	MATCH_JOIN_FAIL                              = 37
	FELLOW_SPECTATOR_JOINED                      = 42
	FELLOW_SPECTATOR_LEFT                        = 43
	ALL_PLAYERS_LOADED                           = 45
	MATCH_START                                  = 46
	MATCH_SCORE_UPDATE                           = 48
	MATCH_TRANSFER_HOST                          = 50
	MATCH_ALL_PLAYERS_LOADED                     = 53
	MATCH_PLAYER_FAILED                          = 57
	MATCH_COMPLETE                               = 58
	MATCH_SKIP                                   = 61
	UNAUTHORIZED                                 = 62
	CHANNEL_JOIN_SUCCESS                         = 64
	CHANNEL_INFO                                 = 65
	CHANNEL_KICK                                 = 66
	CHANNEL_AUTO_JOIN                            = 67
	BEATMAP_INFO_REPLY                           = 69
	PRIVILEGES                                   = 71
	FRIENDS_LIST                                 = 72
	PROTOCOL_VERSION                             = 75
	MAIN_MENU_ICON                               = 76
	MONITOR                                      = 80
	MATCH_PLAYER_SKIPPED                         = 81
	USER_PRESENCE                                = 83
	RESTART                                      = 86
	MATCH_INVITE                                 = 88
	CHANNEL_INFO_END                             = 89
	MATCH_CHANGE_PASSWORD                        = 91
	SILENCE_END                                  = 92
	USER_SILENCED                                = 94
	USER_PRESENCE_SINGLE                         = 95
	USER_PRESENCE_BUNDLE                         = 96
	USER_DM_BLOCKED                              = 100
	TARGET_IS_SILENCED                           = 101
	VERSION_UPDATE_FORCED                        = 102
	SWITCH_SERVER                                = 103
	ACCOUNT_RESTRICTED                           = 104
	RTX                                          = 105
	MATCH_ABORT                                  = 106
	SWITCH_TOURNAMENT_SERVER                     = 107
)

func (p OutgoingPacketID) String() string {
	switch p {
	case USER_ID:
		return "USER_ID"
	case SEND_MESSAGE:
		return "SEND_MESSAGE"
	case PONG:
		return "PONG"
	case HANDLE_IRC_CHANGE_USERNAME:
		return "HANDLE_IRC_CHANGE_USERNAME"
	case HANDLE_IRC_QUIT:
		return "HANDLE_IRC_QUIT"
	case USER_STATS:
		return "USER_STATS"
	case USER_LOGOUT:
		return "USER_LOGOUT"
	case SPECTATOR_JOINED:
		return "SPECTATOR_JOINED"
	case SPECTATOR_LEFT:
		return "SPECTATOR_LEFT"
	case SPECTATE_FRAMES:
		return "SPECTATE_FRAMES"
	case VERSION_UPDATE:
		return "VERSION_UPDATE"
	case SPECTATOR_CANT_SPECTATE:
		return "SPECTATOR_CANT_SPECTATE"
	case GET_ATTENTION:
		return "GET_ATTENTION"
	case NOTIFICATION:
		return "NOTIFICATION"
	case UPDATE_MATCH:
		return "UPDATE_MATCH"
	case NEW_MATCH:
		return "NEW_MATCH"
	case DISPOSE_MATCH:
		return "DISPOSE_MATCH"
	case TOGGLE_BLOCK_NON_FRIEND_DMS:
		return "TOGGLE_BLOCK_NON_FRIEND_DMS"
	case MATCH_JOIN_SUCCESS:
		return "MATCH_JOIN_SUCCESS"
	case MATCH_JOIN_FAIL:
		return "MATCH_JOIN_FAIL"
	case FELLOW_SPECTATOR_JOINED:
		return "FELLOW_SPECTATOR_JOINED"
	case FELLOW_SPECTATOR_LEFT:
		return "FELLOW_SPECTATOR_LEFT"
	case ALL_PLAYERS_LOADED:
		return "ALL_PLAYERS_LOADED"
	case MATCH_START:
		return "MATCH_START"
	case MATCH_SCORE_UPDATE:
		return "MATCH_SCORE_UPDATE"
	case MATCH_TRANSFER_HOST:
		return "MATCH_TRANSFER_HOST"
	case MATCH_ALL_PLAYERS_LOADED:
		return "MATCH_ALL_PLAYERS_LOADED"
	case MATCH_PLAYER_FAILED:
		return "MATCH_PLAYER_FAILED"
	case MATCH_COMPLETE:
		return "MATCH_COMPLETE"
	case MATCH_SKIP:
		return "MATCH_SKIP"
	case UNAUTHORIZED:
		return "UNAUTHORIZED"
	case CHANNEL_JOIN_SUCCESS:
		return "CHANNEL_JOIN_SUCCESS"
	case CHANNEL_INFO:
		return "CHANNEL_INFO"
	case CHANNEL_KICK:
		return "CHANNEL_KICK"
	case CHANNEL_AUTO_JOIN:
		return "CHANNEL_AUTO_JOIN"
	case BEATMAP_INFO_REPLY:
		return "BEATMAP_INFO_REPLY"
	case PRIVILEGES:
		return "PRIVILEGES"
	case FRIENDS_LIST:
		return "FRIENDS_LIST"
	case PROTOCOL_VERSION:
		return "PROTOCOL_VERSION"
	case MAIN_MENU_ICON:
		return "MAIN_MENU_ICON"
	case MONITOR:
		return "MONITOR"
	case MATCH_PLAYER_SKIPPED:
		return "MATCH_PLAYER_SKIPPED"
	case USER_PRESENCE:
		return "USER_PRESENCE"
	case RESTART:
		return "RESTART"
	case MATCH_INVITE:
		return "MATCH_INVITE"
	case CHANNEL_INFO_END:
		return "CHANNEL_INFO_END"
	case MATCH_CHANGE_PASSWORD:
		return "MATCH_CHANGE_PASSWORD"
	case SILENCE_END:
		return "SILENCE_END"
	case USER_SILENCED:
		return "USER_SILENCED"
	case USER_PRESENCE_SINGLE:
		return "USER_PRESENCE_SINGLE"
	case USER_PRESENCE_BUNDLE:
		return "USER_PRESENCE_BUNDLE"
	case USER_DM_BLOCKED:
		return "USER_DM_BLOCKED"
	case TARGET_IS_SILENCED:
		return "TARGET_IS_SILENCED"
	case VERSION_UPDATE_FORCED:
		return "VERSION_UPDATE_FORCED"
	case SWITCH_SERVER:
		return "SWITCH_SERVER"
	case ACCOUNT_RESTRICTED:
		return "ACCOUNT_RESTRICTED"
	case RTX:
		return "RTX"
	case MATCH_ABORT:
		return "MATCH_ABORT"
	case SWITCH_TOURNAMENT_SERVER:
		return "SWITCH_TOURNAMENT_SERVER"
	default:
		return fmt.Sprintf("UNKNOWN (%d)", uint16(p))
	}
}
