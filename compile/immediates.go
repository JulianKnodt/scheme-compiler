package compile

import (
	"strconv"
)

func CharRepresentation(a rune) string {
	return string(a<<8 | charMask)
}

const (
	intBottom = 3
)

func ToCompiled(s string) (int, error) {
	n, err := strconv.Atoi(s)
	return ToRepresentation(n), err
}

func ToRepresentation(n int) int {
	return (n << 2) & max
}
