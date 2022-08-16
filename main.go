package main

import (
	"bytes"
	"fmt"
	"hirasawa/bancho"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type LoginData struct {
	username          string
	passwordHash      string
	osuVersion        string
	utcOffset         int
	displayCity       bool
	privateMessages   bool
	osuPathHash       string
	adapters          string
	adaptersHash      string
	uninstallHash     string
	diskSignatureHash string
}

type LoggedPlayer struct {
	LoginData
	osuToken string
}

func handleBancho(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() != "osu!" && r.Method == "GET" {
		fmt.Fprint(w, "Fuck you")
		return
	}

	if r.Method == "POST" {
		if osu_token := r.Header.Get("Osu-Token"); osu_token != "" {
			// Handle bancho packet
			log.Println("Bancho request sent by", osu_token)

			for {
				p, err := bancho.ReadBanchoPacket(r.Body)
				if err == io.EOF {
					log.Println("End of payload")
					break
				} else if err != nil {
					log.Println("Error parsing awesome packet")
					log.Fatal(err)
				}

				log.Printf("Awesome packet recieved: %#v\n", p)
			}

		} else {
			// It's a login
			log.Println("Bancho login request from", r.UserAgent())
			bodyBytes, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Fatal(err)
			}

			body := string(bodyBytes)
			remainder := strings.Fields(body)

			if len(remainder) != 3 {
				log.Println("Body parsing error")
				return
			}

			nickname := remainder[0]
			password_md5 := remainder[1]
			remainder = strings.Split(remainder[2], "|")

			if len(remainder) != 5 {
				log.Println("Body parsing error")
				return
			}

			osu_version := remainder[0]
			utc_offset, err := strconv.Atoi(remainder[1])

			if err != nil {
				log.Fatal(err)
			}

			display_city := remainder[2] == "1"
			client_hashes := strings.Split(remainder[3], ":")
			pm_private := remainder[4] == "1"

			osu_path_md5 := client_hashes[0]
			adapters_str := client_hashes[1]
			adapters_md5 := client_hashes[2]
			uninstall_md5 := client_hashes[3]
			disk_signature := client_hashes[4]

			loginData := &LoginData{
				nickname,
				password_md5,
				osu_version,
				utc_offset,
				display_city,
				pm_private,
				osu_path_md5,
				adapters_str,
				adapters_md5,
				uninstall_md5,
				disk_signature,
			}

			log.Print(nickname, " trying to log in with password ", password_md5)
			fmt.Printf("%#v\n", loginData)

			w.Header().Add("cho-token", "placeholder")

			payload := bytes.NewBuffer(make([]byte, 0))
			payload.Write(bancho.ProtocolVersion(19))
			payload.Write(bancho.UserID(123123))

			fmt.Printf("Payload: %#v\n", payload)
			payload.WriteTo(w)
		}
	}
}

func main() {
	banchoMux := http.NewServeMux()
	banchoMux.HandleFunc("/", handleBancho)

	banchoServer := &http.Server{
		Addr:    ":49152",
		Handler: banchoMux,
	}

	log.Fatal(banchoServer.ListenAndServe())
}
