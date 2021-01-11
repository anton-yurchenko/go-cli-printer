package printer

import (
	"fmt"
)

type bullet int

const (
	// Action contains an emoji for operations
	Action bullet = iota
	// Success contains a success status emoji
	Success
	// Failure contains an error status emoji
	Failure
	// Completed contains an emoji for successfull application execution
	Completed
	// Warning contains a warning status emoji
	Warning
)

const (
	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
)

const defaultIndent int = 2

// String return a text representation of an emoji.
func (b bullet) String() string {
	return [...]string{"🔹", "✅", "❌", "🎉", "⚠️"}[b]
}

// Printer holds printer configuration
type Printer struct {
	Pointer int
}

// New create new printer
func New() *Printer {
	return new(Printer)
}

// Indent pointer one time to the right right.
// There is a limit of 10 characters, after which you may not move more.
func (p *Printer) Indent() *Printer {
	if p.Pointer <= 10 {
		p.Pointer = p.Pointer + defaultIndent
	}

	return p
}

// Outdent pointer one time to the left.
func (p *Printer) Outdent() *Printer {
	if p.Pointer >= defaultIndent {
		p.Pointer = p.Pointer - defaultIndent
	}

	return p
}

// Print text message in current position.
func (p *Printer) Print(text string) {
	indent := ""
	for i := 0; i < p.Pointer; i++ {
		indent = indent + " "
	}

	fmt.Printf("%v%v\n", indent, text)
}

// PrintWithBullet prints text with an emoji prefix in current position.
func (p *Printer) PrintWithBullet(b bullet, text string) {
	p.Print(fmt.Sprintf("%v %v", b.String(), text))
}
