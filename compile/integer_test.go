package compile

import (
	"strconv"
	"testing"
)

func NumStringRepr(n int) string {
	return strconv.Itoa(n << 2)
}

func TestInteger(t *testing.T) {
	if err := Assert("7", NumStringRepr(7), "basic_int"); err != nil {
		t.Error(err)
	}
	if err := Assert("0", NumStringRepr(0), "zero_int"); err != nil {
		t.Error(err)
	}
	if err := Assert("-3", NumStringRepr(-3), "negative_int"); err != nil {
		t.Error(err)
	}
}
