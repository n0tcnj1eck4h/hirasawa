package server

import (
	"bytes"
	"fmt"
	"hirasawa/bancho/outgoing"
	"hirasawa/bancho/userstore"
	"io"
	"log"
	"net/http"
)

type BanchoServer struct {
	Handlers    *Handlers
	PlayerStore *userstore.PlayerStore
}

func New() (b *BanchoServer) {
	b = &BanchoServer{}
	b.Handlers = NewHandlers()
	b.Handlers.InitDefaultHandlers()
	b.PlayerStore = userstore.New()
	return b
}

func (server *BanchoServer) HandleBanchoLogin(w http.ResponseWriter, r *http.Request) {
	loginData, err := readLoginData(r.Body)
	if err != nil {
		log.Println("Failed to parse login data:", err)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	player, err := server.PlayerStore.Login(loginData)
	if err == userstore.NoSuchUser {
		log.Println("User doesn't exist. Registering...")
		player, err = server.PlayerStore.Register(loginData)
		if err != nil {
			log.Panicln("Failed to register player:", err)
			return
		}

		player, err = server.PlayerStore.Login(loginData)
	}

	if err != nil {
		log.Println("Failed to perform login:", err)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	w.Header().Add("cho-token", player.Session.OsuToken)

	payload := &bytes.Buffer{}
	payload.Write(outgoing.ProtocolVersion(19))
	payload.Write(outgoing.UserID(player.ID))
	payload.WriteTo(w)
}

func (server *BanchoServer) HandleBanchoRequest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Osu-Token")
	player, ok := server.PlayerStore.FromToken(token)

	if !ok {
		log.Println("Failed to get player from token", token)
		w.Write(outgoing.Notification("Nice session token mate"))
		w.Write(outgoing.RestartServer(10))
		return
	}

	ctx := &context{
		Player: player,
		Server: server,
		// Packet: gets set by Dispatch()
	}

	player.Session.PacketQueueLock.Lock()
	defer player.Session.PacketQueueLock.Unlock()

	for {
		if err := server.Handlers.Dispatch(ctx, r.Body); err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error parsing bancho request:", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	b := player.Session.PacketQueue.Bytes()
	player.Session.PacketQueue.Reset()
	if len(b) > 0 {
		log.Println("Replying with bytes", b)
	}
	w.Write(b)
}

func (server *BanchoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() != "osu!" && r.Method == "GET" {
		fmt.Fprint(w, "hihi :3")
		return
	}

	if r.Method == "POST" {
		if osu_token := r.Header.Get("Osu-Token"); osu_token != "" {
			log.Println("Bancho request sent by", osu_token)
			server.HandleBanchoRequest(w, r)
		} else {
			log.Println("Bancho login request from", r.UserAgent())
			server.HandleBanchoLogin(w, r)
		}
	}
}
