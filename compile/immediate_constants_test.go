package compile

import (
	"strconv"
	"testing"
)

func TestConstants(t *testing.T) {
	err := RunTests(
		"#t", "111", "true_const",
		"#f", "47", "false_const",
		"()", "63", "nil_const",
		"#\\1", strconv.Itoa(CharRepresentation(byte('1'))), "char 1",
		"#\\a", strconv.Itoa(CharRepresentation(byte('a'))), "char a",
	)
	if err != nil {
		t.Error(err)
	}
}
