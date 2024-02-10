// Package textcolors provides shortcuts for colorizing text. It respsects the
// informal NO_COLOR (https://no-color.org/) specification, meaning if the
// NO_COLOR environment variable is set to anything other than an empty string,
// colorized output will be disabled.
package textcolors

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/text"
)

var colorEnabled bool

func init() {
	colorEnabled = os.Getenv("NO_COLOR") == ""
}

type Color interface {
	Sprint(a ...interface{}) string
	Sprintf(format string, a ...interface{}) string
}

func FgBlack() Color {
	if colorEnabled {
		return text.FgBlack
	}
	return &noColor{}
}

func FgRed() Color {
	if colorEnabled {
		return text.FgRed
	}
	return &noColor{}
}

func FgGreen() Color {
	if colorEnabled {
		return text.FgGreen
	}
	return &noColor{}
}

func FgYellow() Color {
	if colorEnabled {
		return text.FgYellow
	}
	return &noColor{}
}

func FgBlue() Color {
	if colorEnabled {
		return text.FgBlue
	}
	return &noColor{}
}

func FgMagenta() Color {
	if colorEnabled {
		return text.FgMagenta
	}
	return &noColor{}
}

func FgCyan() Color {
	if colorEnabled {
		return text.FgCyan
	}
	return &noColor{}
}

func FgWhite() Color {
	if colorEnabled {
		return text.FgWhite
	}
	return &noColor{}
}

type noColor struct{}

func (n *noColor) Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func (n *noColor) Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
