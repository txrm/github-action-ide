package formatter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/txrm/github-action-ide/internal/actions"
)

func FormatInputs(inputs map[string]actions.Input) {
	if len(inputs) == 0 {
		fmt.Println("Inputs: None")
		return
	}

	fmt.Println("Inputs:")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for key, input := range inputs {

		// major kludge

		wrappedDesc := WrapText(input.Description, "    ", 40) // block limit
		defaultValue := ""
		if input.Default != nil {
			defaultValue = fmt.Sprintf("(Default: %v)", input.Default)
		}
		fmt.Fprintf(writer, "  %s:\t%s %s\n", KeyColor.Sprintf(key), ValueColor.Sprintf(wrappedDesc), ValueColor.Sprintf(defaultValue))
	}
	writer.Flush()
}
