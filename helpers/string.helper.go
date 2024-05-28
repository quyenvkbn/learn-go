package helpers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ToString(s any) string {
	if s == nil {
		return ""
	}
	str := fmt.Sprintf("%v", s)

	return str
}

func ToJsonString(data any) string {
	jsonData, _ := json.Marshal(data)

	return string(jsonData)
}

func StringToMap(data string) map[string]interface{} {
	var old map[string]interface{}
	_ = json.Unmarshal([]byte(data), &old)

	return old
}

func CompareHashAndPassword(currentPassword string, password string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(password))
	if result == nil {
		return true
	}

	return false
}
