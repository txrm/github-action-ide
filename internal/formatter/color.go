package formatter

import "github.com/fatih/color"

var (
	TitleColor = color.New(color.FgHiCyan, color.Bold) // section titles
	KeyColor   = color.New(color.FgGreen, color.Bold)  // key
	ValueColor = color.New(color.FgWhite)              // values
)
