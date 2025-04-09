package cmd

import (
	"os"

	"github.com/0xAFz/ku/cmd/iaas"
	"github.com/0xAFz/ku/cmd/status"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "Kubar",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(iaas.IaaSCmd)
	rootCmd.AddCommand(status.StateCmd)
}
