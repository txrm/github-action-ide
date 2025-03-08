package cmd

import (
	"fmt"

	"github.com/txrm/github-action-ide/internal/actions"
	"github.com/txrm/github-action-ide/internal/formatter"

	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect [action-name]",
	Short: "Inspects a GitHub Action and expands its details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]

		actionData, err := actions.RetrieveAction(action)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		formatter.FormatActionDetails(actionData)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
