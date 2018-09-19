package compile

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	numberRegex = regexp.MustCompile("^-*\\d+$")
	charRegex   = regexp.MustCompile("^#\\.+$")
	exprRegex   = regexp.MustCompile(`\(\s*(\S+)\s(.+)\)`)
)

func Compile(src string, w io.Writer) (err error) {
	steps := []CompileStep{
		CompileStart,
		CompileImmediate,
		CompileExpr,
		CompileEnd,
	}

	return CompileSequential(steps, w, src)
}

// Writes the compilation of src to w, where w is presumed to be a primitive
func CompilePrimitive(w io.Writer, src string) (err error) {
	switch {
	case src == "#t":
		_, err = fmt.Fprintf(w, movResultTemplate, T)
	case src == "#f":
		_, err = fmt.Fprintf(w, movResultTemplate, F)
	case src == "()":
		_, err = fmt.Fprintf(w, movResultTemplate, Nil)
	case numberRegex.MatchString(src):
		n, err := ToCompiled(src)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(w, movResultTemplate, n)
	case charRegex.MatchString(src):

	default:
		err = fmt.Errorf("Cannot evaluate %s", src)
	}
	return err
}

// parses root level operands
func ParseOperands(src string) []string {
	out := make([]string, 0, 1)
	parensCount := 0
	currIndex := 0
	for i, char := range src {
		switch {
		case char == ' ' && parensCount == 0:
			operand := strings.TrimSpace(src[currIndex:i])
			currIndex = i + 1
			if operand == "" {
				continue
			}
			out = append(out, operand)
		case char == '(':
			parensCount++
		case char == ')':
			parensCount--
		}
	}
	lastOperand := strings.TrimSpace(src[currIndex:])
	if lastOperand != "" {
		out = append(out, lastOperand)
	}
	return out
}

func CompileExpr(w io.Writer, src string) (err error) {
	switch {
	case exprRegex.MatchString(src):
		matches := exprRegex.FindAllStringSubmatch(src, 2)
		if matches == nil {
			return fmt.Errorf("Value matched regex but no submatches")
		}
		body := matches[0]
		operator := body[1]
		operands := ParseOperands(body[2])
		fn, ok := builtin[operator]
		switch {
		case !ok:
			return fmt.Errorf("No such function %s", operator)
		case fn.numArgs != len(operands):
			return fmt.Errorf(
				"Incorrect Number of arguments for %s, expected: %d, instead got list of %s",
				operator, fn.numArgs, strings.Join(operands, ", "))
		default:
			return fn.compile(w, operands...)
		}
	}
	return nil
}

func IsImmediate(src string) bool {
	switch src {
	case TrueString, FalseString, NullString:
		return true
	}

	switch {
	case numberRegex.MatchString(src), charRegex.MatchString(src):
		return true
	}
	return false
}

type CompileStep func(io.Writer, string) error

func CompileStart(w io.Writer, s string) (err error) {
	_, err = fmt.Fprintf(w, startText)
	return err
}

func CompileEnd(w io.Writer, s string) (err error) {
	_, err = fmt.Fprintf(w, "\tret")
	return err
}

func CompileSequential(todo []CompileStep, w io.Writer, src string) (err error) {
	for _, item := range todo {
		err = item(w, src)
		if err != nil {
			return
		}
	}
	return
}

func CompileImmediate(w io.Writer, src string) (err error) {
	if !IsImmediate(src) {
		return nil
	}
	return CompilePrimitive(w, src)
}
