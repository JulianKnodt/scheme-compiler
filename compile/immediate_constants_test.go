package compile

import (
	"testing"
)

func TestConstants(t *testing.T) {
	err := RunTests(
		"#t", "111", "true_const",
		"#f", "47", "false_const",
		"()", "63", "nil_const",
		"#\\1", CharRepresentation('1'), "char 1",
		"#\\a", CharRepresentation('a'), "char 1",
	)
	if err != nil {
		t.Error(err)
	}
}
