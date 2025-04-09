package iaas

import (
	"fmt"
	"time"

	"github.com/0xAFz/ku/internal/api"
	"github.com/0xAFz/ku/internal/state"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Creates infrastructure according to kubar configuration files in the current directory.",
	Run: func(_ *cobra.Command, _ []string) {
		desired, err := state.ReadDesiredState()
		if err != nil {
			fmt.Println(err)
			return
		}
		current, err := state.ReadCurrentState()
		if err != nil {
			fmt.Println(err)
			return
		}

		desiredMap := make(map[string]api.KubarInstanceRequest)
		for _, req := range desired {
			desiredMap[req.Name] = req
		}
		currentMap := make(map[string]api.KubarInstance)
		for _, vm := range current {
			currentMap[vm.Name] = vm
		}

		var newState []api.KubarInstance

		for _, req := range desired {
			if existing, exists := currentMap[req.Name]; !exists || existing.IP != nil {
				// Create VM if it doesn’t exist or isn’t active
				fmt.Printf("kubar_compute_instance.%s: Creating...\n", req.Name)
				if err := provider.CreateInstance(req); err != nil {
					fmt.Printf("%s: %v\n", req.Name, err)
					continue
				}
				start := time.Now()
				waitCount := 1
				for {
					fmt.Printf("kubar_compute_instance.%s: Still creating... [%ds elapsed]\n", req.Name, waitCount)
					time.Sleep(time.Second * 1)
					waitCount++
					ins, err := provider.GetInstance(req.Name)
					if err != nil {
						fmt.Printf("failed to get resource: %v\n", err)
						continue
					}
					if ins.IP == nil {
						continue
					}
					newState = append(newState, *ins)
					break
				}
				fmt.Printf("kubar_compute_instance.%s: Creation complete after %v\n", req.Name, time.Since(start))
			} else {
				// Keep existing VM
				newState = append(newState, existing)
			}
		}

		// Process current state: destroy unwanted VMs
		for name, vm := range currentMap {
			if _, keep := desiredMap[name]; !keep {
				if err := provider.DeleteInstance(map[string]string{"name": vm.Name}); err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("Destroyed: %s\n", name)
			}
		}

		if err := state.WriteCurrentState(newState); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}
