package cmd

import "github.com/spf13/cobra"

var stopCommand = &cobra.Command{
	Use:   "stop",
	Short: "Stop it",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCommand.AddCommand(stopCommand)
}
