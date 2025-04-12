package util

import "encoding/json"

func ToJSONRawMessage(v interface{}) json.RawMessage {
	data, _ := json.Marshal(v)
	return data
}
