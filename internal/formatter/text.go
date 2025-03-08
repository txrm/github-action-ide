package formatter

import (
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

func WrapText(text string, indent string, width uint) string {
	wrapped := wordwrap.WrapString(text, width)
	return indent + strings.ReplaceAll(wrapped, "\n", "\n"+indent)
}
