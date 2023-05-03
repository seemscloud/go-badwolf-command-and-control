package hosts

type Server struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func GetHosts() []Server {
	servers := []Server{
		{
			Hostname: "example.com",
			IP:       "192.168.0.1",
		},
		{
			Hostname: "example.net",
			IP:       "192.168.0.2",
		},
	}

	return servers
}
