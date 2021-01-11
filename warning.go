package printer

import "fmt"

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
