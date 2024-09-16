package models

import (
	"fmt"
	"log"
)

type Monitor struct {
	CamNumber int
	Messages  []Message
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func NewMonitorWithParams(camNumber int, messages []Message) *Monitor {
	return &Monitor{
		CamNumber: camNumber,
		Messages:  messages,
	}
}

func (m *Monitor) SendMessages() {
	try {
		echoClient, err := NewEchoClient(Settings.GetHostMonitor(m.CamNumber))
		if err != nil {
			log.Printf("Failed to create EchoClient: %v", err)
			return
		}
		defer echoClient.Close()

		// Clear
		err = echoClient.SendEchoWithoutReceive([]byte{0x03, 0x44, 0x47})
		if err != nil {
			log.Printf("Failed to send clear message: %v", err)
			return
		}

		for _, item := range m.Messages {
			err = echoClient.SendEchoWithoutReceive(item.Text, item.X, item.Y, item.Color)
			if err != nil {
				log.Printf("Failed to send message: %v", err)
			}
		}
	} catch (err error) {
		log.Printf("Failed to send message. %v", err)
		log.Printf("sendMessages Error: %v", err)
	}
}

type Message struct {
	Text  string
	X     int
	Y     int
	Color string
}

