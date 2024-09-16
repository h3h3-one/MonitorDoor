package models

import (
	"log"

	"your_project/service"
)

var jsonService = service.NewJsonService()

func GetHostDoor(cameraNumber int) *Host {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Error: %v", r)
		}
	}()

	configParam := jsonService.GetConfigParam()
	doorInfo := configParam.MonitorDoorDictionary[cameraNumber]
	ipDoor := doorInfo.IpDoor
	portDoor := doorInfo.PortDoor

	log.Printf("Получение информации о хосте двери. IP: %s PORT: %d", ipDoor, portDoor)
	return &Host{IP: ipDoor, Port: portDoor}
}

func GetHostMonitor(cameraNumber int) *Host {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Не найден ip или port монитора.")
		}
	}()

	configParam := jsonService.GetConfigParam()
	monitorInfo := configParam.MonitorDoorDictionary[cameraNumber]
	ipMonitor := monitorInfo.IpMonitor
	portMonitor := monitorInfo.PortMonitor

	log.Printf("Получение информации о хосте монитора. IP: %s PORT: %d", ipMonitor, portMonitor)
	return &Host{IP: ipMonitor, Port: portMonitor}
}

type Host struct {
	IP   string
	Port int
}
