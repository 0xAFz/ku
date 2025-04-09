package state

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xAFz/ku/internal/api"
)

const (
	current = ".kustate.json"
	desired = "kubar.json"
)

func ReadDesiredState() ([]api.KubarInstanceRequest, error) {
	data, err := os.ReadFile(desired)
	if err != nil {
		return nil, fmt.Errorf("reading desired state: %v", err)
	}
	var reqs []api.KubarInstanceRequest
	err = json.Unmarshal(data, &reqs)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling desired state: %v", err)
	}
	return reqs, nil
}

func ReadCurrentState() ([]api.KubarInstance, error) {
	data, err := os.ReadFile(current)
	if os.IsNotExist(err) {
		return []api.KubarInstance{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading current state: %v", err)
	}
	var state []api.KubarInstance
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling current state: %v", err)
	}
	return state, nil
}

func WriteCurrentState(state []api.KubarInstance) error {
	file, err := json.MarshalIndent(state, "", "    ")
	if err != nil {
		return fmt.Errorf("marshaling current state: %v", err)
	}
	return os.WriteFile(current, file, 0o644)
}
