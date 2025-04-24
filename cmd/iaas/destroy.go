package iaas

import (
	"fmt"
	"sync"

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

		if len(current) == 0 {
			fmt.Println("No objects need to be destroyed.")
			return
		}

		var wg sync.WaitGroup
		var mu sync.Mutex
		var nextState []api.KubarInstance

		for _, v := range current {
			resource := v
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := provider.DeleteInstance(map[string]string{"name": resource.Name}); err != nil {
					fmt.Printf("kubar_compute_instance.%s: %v\n", resource.Name, err)
					mu.Lock()
					nextState = append(nextState, resource)
					mu.Unlock()
					return
				}
				fmt.Printf("kubar_compute_instance.%s: Destruction complete\n", resource.Name)
			}()
		}

		if err := state.WriteCurrentState(nextState); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
