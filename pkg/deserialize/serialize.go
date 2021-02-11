package deserialize

import "encoding/json"

//ToMAP Deserialize string to map
func ToMAP(serialized string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(serialized), &data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}
