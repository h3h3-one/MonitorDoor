package models

type ConfigurationModelDoorAndMonitor struct {
	MqttUsername          string
	MqttPassword          string
	IdClient              string
	IpClient              string
	PortClient            int
	MonitorDoorDictionary map[int]MonitorDoor
}

type MonitorDoor struct {
	Field1 string
	Field2 int
	Field3 string
	Field4 int
}

func NewConfigurationModelDoorAndMonitor() *ConfigurationModelDoorAndMonitor {
	return &ConfigurationModelDoorAndMonitor{
		MqttUsername: "admin",
		MqttPassword: "333",
		IdClient:     "MonitorDoor",
		IpClient:     "194.87.237.67",
		PortClient:   1883,
		MonitorDoorDictionary: map[int]MonitorDoor{
			1: {"1244", 1234, "192.168.8.110", 1985},
			2: {"1244", 1234, "192.168.8.111", 1245},
		},
	}
}
