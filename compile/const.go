package compile

const (
	startText = `
  .text
  .globl _scheme_entry
_scheme_entry:
`

	TrueString        = "#t"
	FalseString       = "#f"
	NullString        = "()"
	charMask          = 0x0F
	movResultTemplate = "\tmovl $%d, %%eax\n"
)

const (
	// represents the value true
	T = 0x6F
	// represents the value false
	F = 0x2F
	// represents the null value "()"
	Nil = 0x3F

	// Mask for 32 bytes
	max = 0xFFFFFFFF
)
