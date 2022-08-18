package bancho

import (
	"bytes"
	"errors"
	"hirasawa/bancho/common"
	"hirasawa/bancho/incoming"
	"hirasawa/bancho/outgoing"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HandleBanchoLogin(w http.ResponseWriter, r *http.Request) {
	loginData, err := readLoginData(r.Body)
	if err != nil {
		log.Println("Failed to parse login data:", err)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	player, err := common.PerformLogin(loginData)
	if err != nil {
		log.Println("Failed to perform login:", err)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	w.Header().Add("cho-token", player.OsuToken)

	payload := &bytes.Buffer{}
	payload.Write(outgoing.ProtocolVersion(19))
	payload.Write(outgoing.UserID(69))
	payload.WriteTo(w)
}

func HandleBanchoRequest(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Osu-Token")
	player, ok := common.GetPlayer(token)

	if !ok {
		log.Println("Failed to get player from token", token)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	context := &common.Context{Player: player}

	player.PacketQueueLock.Lock()
	defer player.PacketQueueLock.Unlock()

	for {
		if p, err := incoming.ReadIncomingBanchoPacket(r.Body); err == nil {
			log.Println("Handling packet", p.Type(), "with size", p.Len())
			p.Handle(context)
		} else if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error parsing bancho request:", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	b := player.PacketQueue.Bytes()
	log.Println("Replying with bytes", b)
	w.Write(b)
}

func readLoginData(r io.Reader) (*common.LoginData, error) {
	bodyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	body := string(bodyBytes)
	remainder := strings.Fields(body)
	if len(remainder) != 3 {
		return nil, errors.New("Login request body is misformatted")
	}

	loginData := &common.LoginData{}
	loginData.Username = remainder[0]
	loginData.PasswordHash = remainder[1]

	remainder = strings.Split(remainder[2], "|")
	if len(remainder) != 5 {
		return nil, errors.New("Login request body is misformatted")
	}

	loginData.OsuVersion = remainder[0]
	loginData.UtcOffset, err = strconv.Atoi(remainder[1])
	if err != nil {
		return nil, errors.New("Error parsing UTC offset")
	}

	loginData.DisplayCity = remainder[2] != "0"
	loginData.PrivateMessages = remainder[4] != "0"

	client_hashes := strings.Split(remainder[3], ":")
	if len(client_hashes) < 5 {
		return nil, errors.New("Misformatted client hashes")
	}

	loginData.OsuPathHash = client_hashes[0]
	loginData.Adapters = client_hashes[1]
	loginData.AdaptersHash = client_hashes[2]
	loginData.UninstallHash = client_hashes[3]
	loginData.DiskSignatureHash = client_hashes[4]

	return loginData, nil
}
