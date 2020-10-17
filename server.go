package control

// Server represents a Server definition as returned by the Servers Control API
type Server struct {
	FQDN        string `json:"fqdn,omitempty"`
	Customer    string `json:"customer,omitempty"`
	Environment string `json:"environment,omitempty"`
	Project     string `json:"project,omitempty"`
	Role        string `json:"role,omitempty"`
	Stage       string `json:"stage,omitempty"`
	Location    string `json:"location,omitempty"`
	Region      string `json:"region,omitempty"`
	ModDate     int    `json:"modDate,omitempty"`
	ModUser     string `json:"modUser,omitempty"`
}
