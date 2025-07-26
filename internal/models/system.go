package models

type SystemInfo struct {
	User     string `json:"user"`
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Kernel   string `json:"kernel"`
	Uptime   string `json:"uptime"`
	Shell    string `json:"shell"`
	CPU      string `json:"cpu"`
	Memory   string `json:"memory"`
	Disk     string `json:"disk"`
	Packages string `json:"packages"`
	GPU      string `json:"gpu"`
}

type InfoItem struct {
	Label string
	Value string
	Color string
}
