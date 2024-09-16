package models

type Host struct {
	IP   string
	Port int
}

func NewHost(ip string, port int) *Host {
	return &Host{
		IP:   ip,
		Port: port,
	}
}
