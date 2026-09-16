package set

import "github.com/spf13/cobra"

var Command = cobra.Command{
	Use: "set",
	Short: "Entre subcommand for task state management",
}

func init() {
	Command.AddCommand(
		&doneCommand,
		&undoneCommand,
	)
}