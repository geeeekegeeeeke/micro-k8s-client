package response

type FirewallBaseInfo struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Version    string `json:"version"`
	PingStatus string `json:"pingStatus"`
}
