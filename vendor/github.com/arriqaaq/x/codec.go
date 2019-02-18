package x

import (
	"encoding/json"
)

func MarshalJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func UnmarshalJson(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
