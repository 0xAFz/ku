package iaas

import (
	"fmt"

	"github.com/0xAFz/ku/internal/api"
	"github.com/0xAFz/ku/internal/state"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy kubar-managed infrastructure.",
	Run: func(_ *cobra.Command, _ []string) {
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range current {
			if err := provider.DeleteInstance(map[string]string{"name": v.Name}); err != nil {
				fmt.Printf("kubar_compute_instance.%s: %v\n", v.Name, err)
				return
			}
			fmt.Printf("kubar_compute_instance.%s: Destruction complete\n", v.Name)
		}

		if err := state.WriteCurrentState([]api.KubarInstance{}); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
