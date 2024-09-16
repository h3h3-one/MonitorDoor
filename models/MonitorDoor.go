package models

type MonitorDoor struct {
	IPDoor      string
	PortDoor    int
	IPMonitor   string
	PortMonitor int
}

func NewMonitorDoor(ipDoor string, portDoor int, ipMonitor string, portMonitor int) *MonitorDoor {
	return &MonitorDoor{
		IPDoor:      ipDoor,
		PortDoor:    portDoor,
		IPMonitor:   ipMonitor,
		PortMonitor: portMonitor,
	}
}
