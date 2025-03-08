package formatter

import (
	"fmt"

	"github.com/txrm/github-action-ide/internal/actions"
)

func FormatRuns(runs actions.Runs) {
	fmt.Println("\nRuns:")
	fmt.Printf("  %s: %s\n", KeyColor.Sprintf("Using"), ValueColor.Sprintf(runs.Using))
	fmt.Printf("  %s: %s\n", KeyColor.Sprintf("Main"), ValueColor.Sprintf(runs.Main))
	if runs.Post != "" {
		fmt.Printf("  %s: %s\n", KeyColor.Sprintf("Post"), ValueColor.Sprintf(runs.Post))
	}
}
