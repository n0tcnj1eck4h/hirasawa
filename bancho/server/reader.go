package server

import (
	"encoding/binary"
	"errors"
	"hirasawa/bancho/userstore"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func readLoginData(r io.Reader) (*userstore.LoginData, error) {
	bodyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	body := string(bodyBytes)
	remainder := strings.Fields(body)
	if len(remainder) != 3 {
		return nil, errors.New("Login request body is misformatted")
	}

	loginData := &userstore.LoginData{}
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

func readInt32List16(r io.Reader) ([]int32, error) {
	var length int16
	err := binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return nil, err
	}

	ids := make([]int32, length)
	err = binary.Read(r, binary.LittleEndian, ids)

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func readString(r io.Reader) (string, error) {
	var ok byte
	binary.Read(r, binary.LittleEndian, &ok)

	if ok != 0x0B {
		return "", nil
	}

	length := byte(0)
	shift := byte(0)

	for {
		var b byte
		err := binary.Read(r, binary.LittleEndian, &b)
		if err != nil {
			return "", err
		}

		length |= b & 0x7F << shift
		if b&0x80 == 0 {
			break
		}

		shift += 7
	}

	ret := make([]byte, length)
	_, err := r.Read(ret)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}
