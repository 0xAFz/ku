package iaas

import (
	"fmt"

	"github.com/0xAFz/ku/internal/api"
	"github.com/0xAFz/ku/internal/config"
	"github.com/0xAFz/ku/internal/vm"
	"github.com/spf13/cobra"
)

var provider *vm.Provider

var IaaSCmd = &cobra.Command{
	Use:   "iaas",
	Short: "Manage iaas actions",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiClient := api.NewAPIClient(vm.BaseURL, config.AppConfig.ApiKey)
		provider = vm.NewProvider(apiClient)
	},
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Usage: kumo iaas <action>")
	},
}

func init() {
	config.LoadConfig()

	IaaSCmd.AddCommand(applyCmd)
	IaaSCmd.AddCommand(destroyCmd)
}
