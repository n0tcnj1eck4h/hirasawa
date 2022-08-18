package common

import "fmt"

type PlayerStore struct {
	PlayerFromID map[int32]*LoggedPlayer
	IdFromToken map[string]int32
}

var playerStore PlayerStore
var idPlaceholder int32 = 10

func init() {
	playerStore.PlayerFromID = map[int32]*LoggedPlayer{}
	playerStore.IdFromToken = map[string]int32{}
}

func PerformLogin(login *LoginData) (*LoggedPlayer, error) {
	id := idPlaceholder
	idPlaceholder++

	player, ok := playerStore.PlayerFromID[id]		
	if ok {
		return player, nil
	} 

	player = &LoggedPlayer{}
	token := fmt.Sprintf("placeholder%d", id)

	player.LoginData = login
	player.ID = id
	player.OsuToken = token

	playerStore.PlayerFromID[id] = player
	playerStore.IdFromToken[token] = id

	return player, nil

}

func GetPlayer(token string) (*LoggedPlayer, bool) {
	id, ok := playerStore.IdFromToken[token]
	if !ok {
		return nil, ok
	}

	player, ok := playerStore.PlayerFromID[id]
	if !ok {
		return nil, ok
	}
	
	return player, ok
}

