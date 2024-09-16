package models

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"

	"golang.org/x/text/encoding/charmap"
)

type EchoClient struct {
	socket  *net.UDPConn
	address *net.UDPAddr
	port    int
	ip      string
}

func NewEchoClient(host Host) *EchoClient {
	client := &EchoClient{
		ip:   host.IP,
		port: host.Port,
	}

	log.Printf("Подключение к клиенту. HOST: %s:%d", client.ip, client.port)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", client.ip, client.port))
	if err != nil {
		log.Printf("EchoClient Ошибка: %v", err)
		return nil
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Printf("EchoClient Ошибка: %v", err)
		return nil
	}

	client.socket = conn
	client.address = addr

	conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	return client
}

func (c *EchoClient) SendEchoWithOutReceive(msg []byte) {
	log.Printf("Попытка отправки сообщения. MESSAGE: %v ADDRESS: %v PORT: %d", msg, c.address, c.port)

	_, err := c.socket.Write(msg)
	if err != nil {
		log.Printf("Сокет не найден.")
		log.Printf("��шибка: %v", err)
	}
}

func BCCCalc(dataBCC []byte, BCCSize int) byte {
	var result byte
	for i := 0; i < BCCSize; i++ {
		result ^= dataBCC[i]
	}
	return result
}

func (c *EchoClient) SendEchoWithoutReceive(message string, x, y, color byte) {
	log.Printf("Отправка сообщения на табло. MESSAGE: %s X: %d Y: %d COLOR: %d", message, x, y, color)

	encoder := charmap.Windows1251.NewEncoder()
	textByte, _ := encoder.Bytes([]byte(message))

	msg := make([]byte, 6+len(textByte))
	msg[0] = byte(len(msg))
	msg[1] = 0x46
	msg[2] = x
	msg[3] = y
	msg[4] = color

	copy(msg[5:], textByte)

	receive := make([]byte, 10)
	_, err := c.socket.Read(receive)
	if err != nil {
		log.Printf("Сокет не найден.")
		return
	}

	log.Printf("Получен ответ от контроллера. Сообщение: %s ADDRESS: %v PORT: %d",
		convertByteToHex(receive), c.address, c.port)

	bcc := BCCCalc(msg, len(msg))
	msg[5+len(textByte)] = bcc

	_, err = c.socket.Write(msg)
	if err != nil {
		log.Printf("Ошибка отправки: %v", err)
		return
	}

	time.Sleep(50 * time.Millisecond)
	log.Printf("Сообщение успешно отправлено. Сообщение: %s", message)
}

func (c *EchoClient) Close() {
	c.socket.Close()
	log.Printf("Сокет закрыт. IP: %s PORT: %d", c.ip, c.port)
}

func convertByteToHex(text []byte) string {
	return hex.EncodeToString(text)
}
