package fuzz

import (
	"unicode/utf8"

	"github.com/timbutler/zxcvbn"
)

func Fuzz(data []byte) int {
	password := string(data)

	_ = zxcvbn.PasswordStrength(password, nil)
	if !utf8.ValidString(password) {
		return 0
	}
	return 1
}
