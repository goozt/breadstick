package main

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

var DEFAULT_ENTROPY = 32

func GenerateSecureToken(length ...int) string {

	byteLength := DEFAULT_ENTROPY
	if len(length) > 0 && length[0] != 0 {
		byteLength = length[0]
	}

	println(byteLength)
	token_byte := make([]byte, byteLength)
	if _, err := rand.Read(token_byte); err != nil {
		return err.Error()
	}

	encoded := base64.StdEncoding.EncodeToString(token_byte)
	encoded = strings.ReplaceAll(encoded, "+", "")
	encoded = strings.ReplaceAll(encoded, "/", "")
	return strings.ReplaceAll(encoded, "=", "")
}
