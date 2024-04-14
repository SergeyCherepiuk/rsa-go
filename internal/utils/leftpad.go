package utils

func LeftPad(message []byte, ch byte, n int) []byte {
	prefix := make([]byte, n)
	for i := range prefix {
		prefix[i] = ch
	}
	return append(prefix, message...)
}
