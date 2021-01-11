package printer

import (
	"fmt"
	"strings"
)

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
