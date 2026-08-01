package encoder

import (
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(counter uint64) (string, error) {
	if counter == 0 {
		return string(alphabet[0]), nil
	}

	var sb strings.Builder

	for counter > 0 {
		reminder := counter % 62
		sb.WriteByte(alphabet[reminder])
		counter = counter / 62
	}

	encoded := sb.String()
	runes := []byte(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
