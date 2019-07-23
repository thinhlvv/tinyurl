package utils

import (
	"strings"
)

func EncodeBase62(deci uint64) string {
	var s string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := strings.Split(s, "")
	var hash_str string
	for deci > 0 {
		hash_str = str[deci%62] + hash_str
		deci /= 62
	}
	return hash_str
}
