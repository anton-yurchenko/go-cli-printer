package printer

import (
	"fmt"
	"strings"
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

func (p *Printer) success(bullet string) {
	p.Print(fmt.Sprintf("%v%v%v%v", bullet, string(colorGreen), "Success", string(colorReset)))
}

// Success prints 'Success' in current position.
func (p *Printer) Success() {
	p.success("")
}

// SuccessWithBullet prints 'Success' with a success emoji prefix in current position.
func (p *Printer) SuccessWithBullet() {
	p.success(fmt.Sprintf("%v ", Success.String()))
}

func (p *Printer) failure(bullet string) {
	p.Print(fmt.Sprintf("%v%v%v%v", bullet, string(colorRed), "Failure", string(colorReset)))
}

// Failure prints 'Failure' in current position.
func (p *Printer) Failure() {
	p.failure("")
}

// FailureWithBullet prints 'Failure' with a failure emoji prefix in current position.
func (p *Printer) FailureWithBullet() {
	p.failure(fmt.Sprintf("%v ", Failure.String()))
}

func (p *Printer) warning(bullet string, text string) {
	p.Print(fmt.Sprintf("%v%v%v%v", bullet, string(colorYellow), text, string(colorReset)))
}

// Warning prints a warning message in current position.
func (p *Printer) Warning(text string) {
	p.warning("", text)
}

// WarningWithBullet prints a warning message with a warning emoji prefix in current position.
func (p *Printer) WarningWithBullet(text string) {
	p.warning(fmt.Sprintf("%v ", Warning.String()), text)
}

// ListOfWarnings prints an slice of message, line by line, as a warnings.
func (p *Printer) ListOfWarnings(text []string) {
	for _, t := range text {
		p.Warning(t)
	}
}

// ListOfWarningsWithBullets prints an slice of message, line by line, as a warnings with warning emoji prefix.
func (p *Printer) ListOfWarningsWithBullets(text []string) {
	for _, t := range text {
		p.WarningWithBullet(t)
	}
}

func (p *Printer) message(title, body string) string {
	var indent string
	for i := 0; i < p.Pointer+defaultIndent; i++ {
		indent = indent + " "
	}

	msg := []string{fmt.Sprintf("%v:", title)}
	msg = append(msg, "")
	for _, l := range strings.Split(body, "\n") {
		msg = append(msg, fmt.Sprintf("%v%v", indent, l))
	}
	msg = append(msg, "")

	return strings.Join(msg, "\n")
}

// WarningMessage prints complex message with a title, all as a warning.
// Example:
//  expected output
//
//  apiVersion: v1
//  kind: ConfigMap
//  metadata:
//    namespace: kube-system
//  data:
//    maxCount: 1
func (p *Printer) WarningMessage(title, body string) {
	p.warning("", p.message(title, body))
}

// WarningMessageWithBullet prints complex message with a title and a warning emoji prefix, all as a warning.
// Example:
//  ⚠️ expected output
//
//  apiVersion: v1
//  kind: ConfigMap
//  metadata:
//    namespace: kube-system
//  data:
//    maxCount: 1
func (p *Printer) WarningMessageWithBullet(title, body string) {
	p.warning(fmt.Sprintf("%v ", Warning.String()), p.message(title, body))
}
