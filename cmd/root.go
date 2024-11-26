package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timwehrle/asars/cmd/config"
)

func init() {
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(TasksCmd)
	rootCmd.AddCommand(TaskCmd)
}

var rootCmd = &cobra.Command{
	Use:     "asars",
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
