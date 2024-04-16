package ascii

import (
	"fmt"
	"strconv"

	"github.com/SergeyCherepiuk/rsa-go/internal/splitter"
	"github.com/SergeyCherepiuk/rsa-go/internal/utils"
)

func Encode(message []byte) []byte {
	codes := make([]byte, 0)
	for _, b := range message {
		code := fmt.Sprintf("%03d", uint8(b))
		codes = append(codes, code...)
	}
	return codes
}

func Decode(message []byte) []byte {
	count := 3 - len(message)%3
	if count == 3 {
		count = 0
	}

	message = utils.LeftPad(message, '0', count)
	codes := splitter.Split(message, 3)

	decoded := make([]byte, len(codes))
	for i, c := range codes {
		code, _ := strconv.ParseUint(string(c), 10, 8)
		decoded[i] = byte(code)
	}

	return decoded
}
