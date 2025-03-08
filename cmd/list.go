package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/txrm/github-action-ide/internal/actions"
	"github.com/txrm/github-action-ide/internal/formatter"
)

// WIP
var listCmd = &cobra.Command{
	Use:   "list [repository]",
	Short: "Lists all actions available in a GitHub repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]

		// Doesn't work yet, need to review returns
		actionsList, err := actions.RetrieveAllActions(repo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		formatter.FormatSection("Available GitHub Actions")
		for _, actionData := range actionsList {
			formatter.FormatActionDetails(actionData)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
