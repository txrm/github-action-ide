package formatter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/txrm/github-action-ide/internal/actions"
)

func FormatOutputs(outputs map[string]actions.Output) {
	if len(outputs) == 0 {
		fmt.Println("Outputs: None")
		return
	}

	fmt.Println("Outputs:")
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for key, output := range outputs {
		wrappedDesc := WrapText(output.Description, "    ", 80)
		fmt.Fprintf(writer, "  %s:\t%s\n", KeyColor.Sprintf(key), ValueColor.Sprintf(wrappedDesc))
	}
	writer.Flush()
}
