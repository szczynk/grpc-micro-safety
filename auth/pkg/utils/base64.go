package utils

import (
	"encoding/base64"
)

func EncodeBase64(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return data
}

func DecodeBase64(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
