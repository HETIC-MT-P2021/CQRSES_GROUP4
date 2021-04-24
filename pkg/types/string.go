package types

import (
	"bytes"
	"encoding/gob"
)

func ToSliceByte(strs []string) []byte{
	buffer := &bytes.Buffer{}
	gob.NewEncoder(buffer).Encode(strs)
	return buffer.Bytes()
}