package compile

import (
	"fmt"
	"io"
)

type BuiltinFunction struct {
	numArgs int
	compile CompileFunctionDefinition
}

type CompileFunctionDefinition func(io.Writer, ...string) error

// guaranteed to have the right number of arguments
var builtin = map[string]BuiltinFunction{
	"fxadd1": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) (err error) {
			err = CompilePrimitive(w, args[0])
			if err != nil {
				return
			}
			_, err = fmt.Fprintf(w, "\taddl $%d, %%eax\n", ToRepresentation(1))
			return
		},
	},
	"fxsub1": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) error {
			err := CompilePrimitive(w, args[0])
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(w, "\tsubl $%d, %%eax\n", ToRepresentation(1))
			return err
		},
	},
	"fxzero?": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) error {
			err := CompilePrimitive(w, args[0])
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(w, `
  test %%eax, %%eax
  jz iszero
  movl $%d, %%eax
  ret
iszero:
  movl $%d, %%eax
`, F, T)
			return err
		},
	},
	"boolean?": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) error {
			err := CompilePrimitive(w, args[0])
			if err != nil {
				return nil
			}
			_, err = fmt.Fprintf(w,
				`
  cmp $%d, %%eax
  je isbool
  cmp $%d, %%eax
  je isbool
  movl $%d, %%eax
  ret
isbool:
  movl $%d, %%eax # mov true
`, F, T, F, T)
			return err
		},
	},
	"null?": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) error {
			return nil
		},
	},
	"not": BuiltinFunction{
		numArgs: 1,
		compile: func(w io.Writer, args ...string) error {
			return nil
		},
	},
}
