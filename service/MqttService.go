package service

import (
	"encoding/json"
	"log"
	"net"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	jsonService JsonService
}

func (ms *MqttService) connectedMqtt(host string, port int) {
	log.Printf("Создание подключения клиента... HOST_NAME = %s, PORT = %d, USERNAME = %s, PASSWORD = %s",
		ms.jsonService.getConfigParam().IpClient,
		ms.jsonService.getConfigParam().PortClient,
		ms.jsonService.getConfigParam().MqttUsername,
		ms.jsonService.getConfigParam().MqttPassword)

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://" + host + ":" + string(port)).
		SetClientID(net.JoinHostPort("", "Monitor")).
		SetAutoReconnect(true).
		SetConnectTimeout(5 * time.Second)

	opts.SetUsername(ms.jsonService.getConfigParam().MqttUsername)
	opts.SetPassword(ms.jsonService.getConfigParam().MqttPassword)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка: %v", token.Error())
	}

	ms.subscribe(client)
}

func (ms *MqttService) subscribe(client mqtt.Client) {
	log.Println("Выполнение подписки на топик... ТОПИК: Parking/MonitorDoor/#")
	if token := client.Subscribe("Parking/MonitorDoor/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		ms.handleMessage(msg)
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Ошибка: %v", token.Error())
	}
	log.Println("Подписка на топик Parking/MonitorDoor/# произошла успешно.")
}

func (ms *MqttService) handleMessage(msg mqtt.Message) {
	log.Printf("Получено сообщение! ТОПИК: %s СООБЩЕНИЕ: %s", msg.Topic(), msg.Payload())
	var doors Doors
	jsonStr := string(msg.Payload())

	switch msg.Topic() {
	case "Parking/MonitorDoor/Monitor/View":
		log.Println("Принят топик - Parking/MonitorDoor/Monitor/View")
		var monitor Monitor
		if err := json.Unmarshal(msg.Payload(), &monitor); err == nil {
			monitor.sendMessages()
		} else {
			log.Printf("Ошибка: %v", err)
		}
	case "Parking/MonitorDoor/Doors/Open/0":
		log.Println("Принят топик - Parking/MonitorDoor/Doors/Open/0")
		if err := json.Unmarshal(msg.Payload(), &doors); err == nil {
			doors.openDoor0()
		} else {
			log.Printf("Ошибка: %v", err)
		}
	case "Parking/MonitorDoor/Doors/Open/1":
		log.Println("Принят топик - Parking/MonitorDoor/Doors/Open/1")
		if err := json.Unmarshal(msg.Payload(), &doors); err == nil {
			doors.openDoor1()
		} else {
			log.Printf("Ошибка: %v", err)
		}
	case "Parking/MonitorDoor/1":
		log.Println("Принят топик - Parking/MonitorDoor/1")
		if err := json.Unmarshal([]byte("{\"cameraNumber\": \"4\"}"), &doors); err == nil {
			doors.openDoor1()
		} else {
			log.Printf("Ошибка: %v", err)
		}
	case "Parking/MonitorDoor/0":
		log.Println("Принят топик - Parking/MonitorDoor/0")
		if err := json.Unmarshal([]byte("{\"cameraNumber\": \"2\"}"), &doors); err == nil {
			doors.openDoor0()
		} else {
			log.Printf("Ошибка: %v", err)
		}
	}
}
