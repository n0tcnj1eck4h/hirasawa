package bancho

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	OsuToken string
}

func HandleBanchoLogin(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}

	body := string(bodyBytes)
	remainder := strings.Fields(body)
	if len(remainder) != 3 {
		log.Println("Login request body is misformatted, request ignored...")
		return
	}

	loginData := &LoginData{}
	loginData.Username = remainder[0]
	loginData.PasswordHash = remainder[1]

	remainder = strings.Split(remainder[2], "|")
	if len(remainder) != 5 {
		log.Println("Login request body is misformatted, request ignored...")
		return
	}

	loginData.OsuVersion = remainder[0]
	loginData.UtcOffset, err = strconv.Atoi(remainder[1])
	if err != nil {
		log.Println("Error parsing UTC offset:", err)
		return
	}

	loginData.DisplayCity = remainder[2] != "0"
	loginData.PrivateMessages = remainder[4] != "0"

	client_hashes := strings.Split(remainder[3], ":")
	if len(client_hashes) < 5 {
		log.Println("Invalid client hashes")
		return
	}

	loginData.OsuPathHash = client_hashes[0]
	loginData.Adapters = client_hashes[1]
	loginData.AdaptersHash = client_hashes[2]
	loginData.UninstallHash = client_hashes[3]
	loginData.DiskSignatureHash = client_hashes[4]

	w.Header().Add("cho-token", "placeholder")

	payload := &bytes.Buffer{}
	payload.Write(ProtocolVersion(19))
	payload.Write(UserID(123123))
	payload.WriteTo(w)
}

func HandleBanchoRequest(w http.ResponseWriter, r *http.Request) {
	for {
		if p, err := ReadBanchoPacket(r.Body); err == nil {
			log.Printf("New bancho request: %v, Length: %v\n", p.ReadType, p.Length)
		} else if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error parsing bacho request:", err)
			return
		}
	}

}
