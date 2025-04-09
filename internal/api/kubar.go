package api

type KubarInstanceRequest struct {
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	Flavor    string   `json:"flavor"`
	KeyName   []string `json:"key_name"`
	PublicKey []string `json:"public_key"`
	DiskSize  int      `json:"disk_size"`
}

type KubarInstance struct {
	Name      string  `json:"name"`
	Status    int     `json:"status"`
	Flavor    string  `json:"flavor"`
	DiskSize  int     `json:"disk"`
	StartDate string  `json:"start_date"`
	IP        *string `json:"ip"`
	Image     string  `json:"image"`
	CPU       int     `json:"cpu"`
	Memory    int     `json:"memory"`
}
