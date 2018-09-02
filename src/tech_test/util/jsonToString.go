package util

import "encoding/json"

func JsonToString(payload interface{}) string {
	data, _ := json.Marshal(payload)
	return string(data[:])
}
