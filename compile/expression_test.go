package compile

import (
	"strconv"
	"testing"
)

func TestExpression(t *testing.T) {
	err := RunTests(
		"(fxadd1 1)", NumStringRepr(2), "fxadd1_basic",
		"(fxsub1 1)", NumStringRepr(0), "fxsub1_basic",
		"(fxzero? 1)", strconv.Itoa(F), "fxzero?_1",
		"(fxzero? 0)", strconv.Itoa(T), "fxzero?_0",
		"(fxsub1  0)", NumStringRepr(-1), "fxsub1_negative",
		"( boolean? #f)", NumStringRepr(1), "boolean?_f",
		"( boolean? #t)", NumStringRepr(1), "boolean?_t",
		"( boolean? 1)", NumStringRepr(1), "boolean?_1",
		"( boolean?    #\\3)", NumStringRepr(1), "boolean_char",
		"( null? 1)", strconv.Itoa(F), "null?_1",
	)
	if err != nil {
		t.Error(err)
	}
}
