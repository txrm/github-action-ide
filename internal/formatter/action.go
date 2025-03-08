package formatter

import (
	"fmt"

	"github.com/txrm/github-action-ide/internal/actions"
)

func FormatActionDetails(actionData *actions.GithubAction) {
	FormatSection("GitHub Action Details")
	fmt.Printf("%s: %s\n", KeyColor.Sprintf("Name"), ValueColor.Sprintf(actionData.Name))
	fmt.Printf("%s: %s\n", KeyColor.Sprintf("Description"), ValueColor.Sprintf(WrapText(actionData.Description, "  ", 80)))

	FormatInputs(actionData.Inputs)
	FormatOutputs(actionData.Outputs)
	FormatRuns(actionData.Runs)

	fmt.Println("\nDone!")
}
