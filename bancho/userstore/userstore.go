package userstore

import (
	"errors"
	"fmt"
	"strings"
)

type PlayerStore struct {
	playerMap  map[int32]*Player
	idTokenMap map[string]int32
	idNameMap  map[string]int32
}

var idPlaceholder int32 = 10
var Store PlayerStore

var WrongPassword = errors.New("Wrong password")
var NoSuchUser = errors.New("User doesn't exist")

func init() {
	Store = PlayerStore{}
	Store.playerMap = map[int32]*Player{}
	Store.idTokenMap = map[string]int32{}
	Store.idNameMap = map[string]int32{}
}

func getSafeName(name string) string {
	name = strings.ToLower(name)
	name = strings.Replace(name, " ", "_", -1)
	return name
}

func (store *PlayerStore) Register(login *LoginData) (*Player, error) {
	_, ok := store.FromName(login.Username)
	if ok {
		return nil, errors.New("Username already registered")
	}

	player := &Player{}
	player.ID = idPlaceholder
	player.UsernameSafe = getSafeName(login.Username)
	player.DisplayName = login.Username
	player.PasswordHash = login.PasswordHash
	player.Country = [2]byte{'P', 'L'}
	player.Stats = &PlayerStats{}

	store.Add(player)

	idPlaceholder++
	return player, nil
}

func (store *PlayerStore) Login(login *LoginData) (*Player, error) {
	player, ok := store.FromName(login.Username)

	if !ok {
		return nil, NoSuchUser
	}

	if login.PasswordHash != player.PasswordHash {
		return nil, WrongPassword
	}

	if player.Online() {
		return player, errors.New("Player already online")
	}

	player.Session = &Session{
		OsuToken:  fmt.Sprintf("placeholdertoken%d", player.ID),
		Status: &PlayerStatus{},
		LoginData: login,
	}

	store.idTokenMap[player.Session.OsuToken] = player.ID

	return player, nil
}

func (store *PlayerStore) Add(p *Player) error {
	_, exists := store.FromID(p.ID)
	if exists {
		return errors.New("Player already in memory")
	}

	store.playerMap[p.ID] = p
	store.idNameMap[p.UsernameSafe] = p.ID
	if p.Online() {
		store.idTokenMap[p.Session.OsuToken] = p.ID
	}

	return nil
}

func (store *PlayerStore) Remove(p *Player) {
	delete(store.playerMap, p.ID)
	delete(store.idNameMap, p.UsernameSafe)
	delete(store.idTokenMap, p.Session.OsuToken)
}

func (store *PlayerStore) FromName(name string) (*Player, bool) {
	name_safe := getSafeName(name)

	id, ok := store.idNameMap[name_safe]
	if !ok {
		return nil, false
	}

	player, ok := store.playerMap[id]
	if !ok {
		return nil, false
	}

	return player, true
}

func (store *PlayerStore) FromToken(token string) (*Player, bool) {
	id, ok := store.idTokenMap[token]
	if !ok {
		return nil, ok
	}

	player, ok := store.FromID(id)
	if !ok {
		return nil, ok
	}

	return player, ok
}

func (store *PlayerStore) FromID(id int32) (*Player, bool) {
	player, ok := store.playerMap[id]
	if !ok {
		return nil, ok
	}

	return player, ok
}
