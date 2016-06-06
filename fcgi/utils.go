package fcgi

import (
	"bytes"
	"encoding/binary"
	"io"
)

func readFcgiLength(buf *bytes.Buffer, forceFirstByteEofErr bool) (int, bool, error) {
	c, err := buf.ReadByte()
	if err == io.EOF {
		if forceFirstByteEofErr {
			return 0, false, err
		} else {
			return 0, true, nil
		}
	} else if err != nil {
		return 0, false, err
	}
	keyLength := 0
	if c > 127 {
		l := make([]byte, 4)
		_, err := buf.Read(l[1:])
		if err != nil {
			return 0, false, err
		}
		l[0] = c - 127
		keyLength = int(binary.BigEndian.Uint32(l))
	} else {
		keyLength = int(c) & 0xFF
	}
	return keyLength, false, nil
}
