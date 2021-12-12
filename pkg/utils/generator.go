package utils

import (
	"strconv"
	"strings"
	"time"
)

const (
	Alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Lenght   int    = 13
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func LinkHash() string {
	newkey := make([]string, Lenght)

	z := strconv.FormatInt(time.Now().UnixMilli(), 10)
	z = reverse(z)

	for i := 0; i < Lenght; i++ {
		newkey[i] = string(Alphabet[z[i]])
	}

	return strings.Join(newkey, "")
}
