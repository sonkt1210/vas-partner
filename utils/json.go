package utils

import (
	jsoniter "github.com/json-iterator/go"
)

var jsoniterAPI = jsoniter.ConfigCompatibleWithStandardLibrary

// Marshal marshal an interface
func Marshal(v interface{}) ([]byte, error) {
	return jsoniterAPI.Marshal(v)
}

// Unmarshal unmarshal a byte array to an interface
func Unmarshal(data []byte, v interface{}) error {
	return jsoniterAPI.Unmarshal(data, v)
}
