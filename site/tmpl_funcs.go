package site

import (
	"encoding/json"
)

func stringyfyJson(object any) string {
	stringObject, _ := json.Marshal(object)
	return string(stringObject)
}
