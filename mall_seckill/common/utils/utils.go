package utils

import (
	"encoding/json"
	"fmt"
)

func MapToStr(body map[string]interface{}) string {
	byteBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteBody)
}

func StrToMap(body string) map[string]interface{} {
	var data map[string]interface{}
	_ = json.Unmarshal([]byte(body), &data)
	return data
}
