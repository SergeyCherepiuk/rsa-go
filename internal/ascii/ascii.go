package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

const Separator = " "

func Encode(message []byte) []byte {
	strCodes := make([]string, 0)
	for _, b := range message {
		code := fmt.Sprintf("%d", uint8(b))
		strCodes = append(strCodes, code)
	}
	return []byte(strings.Join(strCodes, Separator))
}

func Decode(message []byte) []byte {
	strCodes := strings.Split(string(message), Separator)

	var decoded string
	for _, strCode := range strCodes {
		code, _ := strconv.ParseUint(strCode, 10, 8)
		decoded += string(byte(code))
	}

	return []byte(decoded)
}
