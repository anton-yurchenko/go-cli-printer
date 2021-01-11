package printer

import "fmt"

func (p *Printer) success(bullet, text string) {
	p.Print(fmt.Sprintf("%v%v%v%v", bullet, string(colorGreen), text, string(colorReset)))
}

// Success prints 'Success' in current position.
func (p *Printer) Success() {
	p.success("", "Success")
}

// SuccessWithBullet prints 'Success' with a success emoji prefix in current position.
func (p *Printer) SuccessWithBullet() {
	p.success(fmt.Sprintf("%v ", Success.String()), "Success")
}

// SuccessMessage prints custom success message in current position.
func (p *Printer) SuccessMessage(text string) {
	p.success("", text)
}

// SuccessMessageWithBullet prints custom success message with a success emoji prefix in current position.
func (p *Printer) SuccessMessageWithBullet(text string) {
	p.success(fmt.Sprintf("%v ", Success.String()), text)
}
