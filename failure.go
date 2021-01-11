package printer

import (
	"fmt"
	"os"
)

func (p *Printer) failure(bullet, text string) {
	p.Print(fmt.Sprintf("%v%v%v%v", bullet, string(colorRed), text, string(colorReset)))
}

// Failure prints 'Failure' in current position.
func (p *Printer) Failure() {
	p.failure("", "Failure")
}

// FailureWithBullet prints 'Failure' with a failure emoji prefix in current position.
func (p *Printer) FailureWithBullet() {
	p.failure(fmt.Sprintf("%v ", Failure.String()), "Failure")
}

// FailureMessage prints custom failure message in current position.
func (p *Printer) FailureMessage(text string) {
	p.failure("", text)
}

// FailureMessageWithBullet prints custom failure message with a success emoji prefix in current position.
func (p *Printer) FailureMessageWithBullet(text string) {
	p.failure(fmt.Sprintf("%v ", Failure.String()), text)
}

// Fatal prints custom failure message in current position and exits execution with exit code 1.
func (p *Printer) Fatal(text string) {
	p.failure("", text)
	os.Exit(1)
}

// FatalWithBullet prints custom failure message with a success emoji prefix in current position and exits execution with exit code 1.
func (p *Printer) FatalWithBullet(text string) {
	p.failure(fmt.Sprintf("%v ", Failure.String()), text)
	os.Exit(1)
}
