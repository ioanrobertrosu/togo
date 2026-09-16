package cli

import (
	"github.com/ioanrobertrosu/todo/internal/cli/set"
	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use: "togo",
	Short: "Program entry command",

	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCommand.Execute()
}

func init() {
	rootCommand.AddCommand(
		&createCommand,
		&editCommand,
		&deleteCommand,
		&listCommand,
		&set.Command,
	)
}