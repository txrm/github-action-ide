package formatter

import (
	"fmt"
	"strings"
)

func FormatSection(title string) {
	fmt.Println()
	TitleColor.Println(title)
	fmt.Println(strings.Repeat("─", 50))
}
