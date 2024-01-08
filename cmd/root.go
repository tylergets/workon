package cmd

import (
	"fmt"
	"github.com/tylergets/workon/config"
	"github.com/tylergets/workon/workflows"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "workon <flowName>",
	Short: "Workon runs the workflow specified by flowName",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flowName := args[0]
		cfg := config.LoadConfig()
		workflows.ExecuteWorkflow(flowName, cfg)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
