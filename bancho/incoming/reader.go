package incoming

import (
	"encoding/binary"
	"io"
)

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
