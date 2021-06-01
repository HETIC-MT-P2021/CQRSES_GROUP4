package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// StringToSliceByte []string => []byte
func StringToSliceByte(strs []string) []byte{
	buffer := &bytes.Buffer{}
	gob.NewEncoder(buffer).Encode(strs)
	return buffer.Bytes()
}

// StringToMAP string => map
func StringToMAP(str string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}