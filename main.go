package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

var logger *log.Logger
var jsonService JsonService
var mqttService MqttService

func init() {
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	jsonService = NewJsonService()
	mqttService = NewMqttService()
}

func main() {
	logger.Println("Starting program.")
	try {
		err := mqttService.ConnectMqtt(
			jsonService.GetConfigParam().IpClient,
			jsonService.GetConfigParam().PortClient,
		)
		if err != nil {
			logger.Printf("Error: %v", err)
			return
		}
		// startBackgroundMethods()
	} catch (err error) {
		logger.Printf("Error: %v", err)
	}
}

/*
func startBackgroundMethods() {
	go func() { // Output time to monitor
		dateFormat := "15:04"
		for {
			messages := []Message{
				{Type: 9, SubType: 15, Command: 1, Data: time.Now().Format(dateFormat)},
			}
			monitor := NewMonitor(1, messages)
			monitor.SendMessages()
			time.Sleep(60 * time.Second)
		}
	}()
}
*/

