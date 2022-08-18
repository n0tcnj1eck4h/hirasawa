package incoming

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

func ReadIncomingBanchoPacket(r io.Reader) (BanchoHandlable, error) {
	var header PacketHeader

	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	switch header.ReadType {
	case PING:
		return Ping{header}, err
	case REQUEST_STATUS_UPDATE:
		return RequestStatusUpdate{header}, err
	case USER_STATS_REQUEST:
		return UserStatsRequest{header, readInt32List16(r)}, err
	default:
		_, err := io.CopyN(ioutil.Discard, r, int64(header.Length))
		if err != nil {
			return nil, err
		}

		return nil, errors.New(fmt.Sprintf("Unsupported packet: %v", header.ReadType))
	}
}

func readInt32List16(r io.Reader) []int32 {
	var length int16
	binary.Read(r, binary.LittleEndian, &length)

	ids := make([]int32, length)
	binary.Read(r, binary.LittleEndian, ids)

	return ids
}
