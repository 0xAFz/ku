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

		var wg sync.WaitGroup
		removeChan := make(chan int, len(current))

		for i := range current {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				if err := provider.DeleteInstance(map[string]string{"name": current[i].Name}); err != nil {
					fmt.Printf("kubar_compute_instance.%s: %v\n", current[i].Name, err)
					return
				}
				fmt.Printf("kubar_compute_instance.%s: Destruction complete\n", current[i].Name)
			}(i)
		}

		var removeIndexes []int
		go func() {
			wg.Wait()
			close(removeChan)
		}()

		for i := range removeChan {
			removeIndexes = append(removeIndexes, i)
		}

		for i := len(removeIndexes) - 1; i >= 0; i-- {
			current = removeResource(current, removeIndexes[i])
		}

		if err := state.WriteCurrentState(current); err != nil {
			fmt.Println("update current state:", err)
			return
		}
	},
}

func removeResource(s []api.KubarInstance, i int) []api.KubarInstance {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
