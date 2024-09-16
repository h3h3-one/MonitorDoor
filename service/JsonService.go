package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigurationModelDoorAndMonitor struct {
}

type JsonService struct{}

func (js *JsonService) isNewFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0644)
	if err != nil {
		if os.IsExist(err) {
			return
		}
		log.Fatalf("Ошибка: %v", err)
	}
	defer file.Close()

	model := ConfigurationModelDoorAndMonitor{}
	jsonData, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	log.Printf("Файл конфигурации успешно создан. Запустите программу заново. ПУТЬ: %s", filePath)
	os.Exit(0)
}

func (js *JsonService) getConfigParam() *ConfigurationModelDoorAndMonitor {
	model := &ConfigurationModelDoorAndMonitor{}
	filePath := "MonitorDoorConfig.json"

	js.isNewFile(filePath)

	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = json.Unmarshal(jsonFile, model)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	return model
}
