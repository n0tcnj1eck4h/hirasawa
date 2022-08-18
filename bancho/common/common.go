package common

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

type LoggedPlayer struct {
	LoginData
	OsuToken        string
	PacketQueue     bytes.Buffer
	PacketQueueLock sync.Mutex
}

type Context struct {
	LoggedPlayer
}
