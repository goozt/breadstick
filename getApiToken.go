package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return err.Error()
	}
	return hex.EncodeToString(b)
}

func main() {
	length := 18
	if len(os.Args) > 1 {
		var err error
		length, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(GenerateSecureToken(length))
}
