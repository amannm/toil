package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "toil",
	Short: "",
	Long:  "",
}

func Execute(ctx context.Context) error {
	return rootCommand.ExecuteContext(ctx)
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
