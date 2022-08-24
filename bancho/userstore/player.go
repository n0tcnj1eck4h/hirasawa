package userstore

import (
	"bytes"
	"sync"
)

type LoginData struct {
	Username          string
	PasswordHash      string
	OsuVersion        string
	UtcOffset         int
	DisplayCity       bool
	PrivateMessages   bool
	OsuPathHash       string
	Adapters          string
	AdaptersHash      string
	UninstallHash     string
	DiskSignatureHash string
}

type PlayerStats struct {
	Mode        byte
	TotalScore  int64
	RankedScore int64
	PP          uint32
	PlayCount   int32
	PlayTime    uint32
	Accuracy    float32
	MaxCombo    uint32
	TotalHits   uint32
	ReplayViews uint32
	// count ranks...
}

type PlayerStatus struct {
	Action   uint8
	InfoText string
	MapHash  string
	Mods     uint32
	Mode     uint8
	MapID    int32
	PrivateMessages   bool
	SilenceEnd	int32
}

type Player struct {
	ID           int32
	UsernameSafe string
	DisplayName  string
	PasswordHash string
	Country      [2]byte
	Session      *Session
	Stats        *PlayerStats
}

type Session struct {
	OsuToken        string
	LoginData       *LoginData
	Status          PlayerStatus
	PacketQueue     bytes.Buffer
	PacketQueueLock sync.Mutex
}

func (p Player) Online() bool {
	return p.Session != nil
}
