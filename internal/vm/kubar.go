package vm

import (
	"encoding/json"
	"fmt"

	"github.com/0xAFz/ku/internal/api"
)

const (
	BaseURL = "https://mz.kubarcloud.com/api/vm"
)

type Provider struct {
	client *api.APIClient
}

func NewProvider(client *api.APIClient) *Provider {
	return &Provider{
		client: client,
	}
}

func (a *Provider) CreateInstance(req api.KubarInstanceRequest) error {
	endpoint := "/create"
	_, err := a.client.Post(endpoint, req)
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}
	return nil
}

func (a *Provider) GetInstance(name string) (*api.KubarInstance, error) {
	endpoint := fmt.Sprintf("/list?name=%s", name)
	resp, err := a.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("get instance: %w", err)
	}
	var instance api.KubarInstance
	if err := json.Unmarshal(resp, &instance); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}
	return &instance, nil
}

func (a *Provider) DeleteInstance(payload map[string]string) error {
	endpoint := "/delete"
	_, err := a.client.Delete(endpoint, payload)
	if err != nil {
		return fmt.Errorf("delete instance: %w", err)
	}
	return nil
}
