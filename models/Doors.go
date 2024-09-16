package models

import (
	"fmt"
	"log"
)

type Doors struct {
	CameraNumber  int
	DataOpenDoor0 []byte
	DataOpenDoor1 []byte
}

func NewDoors() *Doors {
	return &Doors{
		DataOpenDoor0: []byte{0x02, 0x1F, 0x00, 0x03, 0x61, 0x38},
		DataOpenDoor1: []byte{0x02, 0x1F, 0x01, 0x03, 0xb9, 0x21},
	}
}

func (d *Doors) OpenDoor1() {
	d.openDoor(d.DataOpenDoor1)
}

func (d *Doors) OpenDoor0() {
	d.openDoor(d.DataOpenDoor0)
}

func (d *Doors) openDoor(data []byte) {
	try {
		echoClient, err := NewEchoClient(Settings.GetHostDoor(d.CameraNumber))
		if err != nil {
			log.Printf("Error creating EchoClient: %v", err)
			return
		}
		defer echoClient.Close()

		err = echoClient.SendEchoWithOutReceive(data)
		if err != nil {
			log.Printf("Error sending echo: %v", err)
		}
	} catch (err error) {
		log.Printf("Error: %v", err)
	}
}

