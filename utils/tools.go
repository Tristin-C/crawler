package utils

import "encoding/json"

func ToJson(obj interface{}) string {
	str, _ := json.Marshal(obj)
	return string(str)
}
