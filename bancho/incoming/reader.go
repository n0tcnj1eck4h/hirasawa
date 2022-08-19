package incoming

import (
	"encoding/binary"
	"io"
)

func readInt32List16(r io.Reader) []int32 {
	var length int16
	binary.Read(r, binary.LittleEndian, &length)

	ids := make([]int32, length)
	binary.Read(r, binary.LittleEndian, ids)

	return ids
}
