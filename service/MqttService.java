package service;

import com.fasterxml.jackson.databind.ObjectMapper;
import models.Doors;
import models.Monitor;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.eclipse.paho.client.mqttv3.*;

import java.net.InetAddress;

public class MqttService {
    private static final Logger log = LogManager.getLogger(MqttService.class);
    JsonService jsonService = new JsonService();
    static ObjectMapper mapper = new ObjectMapper();

    public void connectedMqtt(String host, int port) {
        log.info(
                "Создание подключения клиента... HOST_NAME = " + jsonService.getConfigParam().getIpClient() +
                        ", PORT = " + jsonService.getConfigParam().getPortClient() +
                        ", USERNAME = " + jsonService.getConfigParam().getMqttUsername() +
                        ", PASSWORD = " + jsonService.getConfigParam().getMqttPassword()
        );
        try {
            MqttClient mqttClient = new MqttClient(
                    "tcp://" + host + ":" + port,
                    InetAddress.getLocalHost() + "-Monitor");
            mqttClient.setCallback(new MqttCallbackExtended() {
                @Override
                public void connectComplete(boolean reconnect, String serverURI) {

                    if(reconnect) {
                        subscribe(mqttClient);
                    }
                }

                @Override
                public void connectionLost(Throwable throwable) {
                    log.warn("Соединение с MQTT потеряно!");
                }

                @Override
                public void messageArrived(String s, MqttMessage mqttMessage) {
                    log.info("Получено сообщение: " + s);
                }

                @Override
                public void deliveryComplete(IMqttDeliveryToken iMqttDeliveryToken) {
                }
            });

                MqttConnectOptions options = new MqttConnectOptions();
                options.setAutomaticReconnect(true);
                options.setConnectionTimeout(5000);
                options.setUserName(jsonService.getConfigParam().getMqttUsername());
                options.setPassword(jsonService.getConfigParam().getMqttPassword().toCharArray());

                log.info(
                        "Выставленные настройки MQTT: " +
                                "Автоматический реконнект = " + options.isAutomaticReconnect()
                );
                mqttClient.connect(options);

                subscribe(mqttClient);
        } catch (Exception e) {
            log.error("Ошибка: " + e);
        }
    }

    private static void subscribe(MqttClient mqttClient) {
        try {
//            log.info("Успешное подключение к MQTT. " + serverURI);
            log.info("Выполнение подписки на топик... ТОПИК: Parking/MonitorDoor/#");
            mqttClient.subscribe("Parking/MonitorDoor/#", (topic, message) -> {
                try {
                    log.info("Получено сообщение! ТОПИК: " + topic + " СООБЩЕНИЕ: " + message);
                    String json = new String(message.getPayload());
                    switch (topic) {
                        case "Parking/MonitorDoor/Monitor/View" -> {
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/Monitor/View");
                                Monitor monitor = mapper.readValue(json, Monitor.class); // Преобразуем JSON в java объект
                                monitor.sendMessages();

                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                        case "Parking/MonitorDoor/Doors/Open/0" -> {
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/Doors/Open/0");
                                var doors = mapper.readValue(json, Doors.class);
                                doors.openDoor0();
                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                        case "Parking/MonitorDoor/Doors/Open/1" -> {
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/Doors/Open/1");
                                var doors = mapper.readValue(json, Doors.class);
                                doors.openDoor1();
                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                        case "Parking/MonitorDoor/Doors/Warning/" -> {
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/Doors/Warning/");
                                //
                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                        case "Parking/MonitorDoor/1" -> { // Тест
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/1");
                                var doors = mapper.readValue("{\"cameraNumber\" : \"4\"}", Doors.class);
                                doors.openDoor1();
                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                        case "Parking/MonitorDoor/0" -> { // Тест
                            try {
                                log.info("Принят топик - Parking/MonitorDoor/0");
                                var doors = mapper.readValue("{\"cameraNumber\" : \"2\"}", Doors.class);
                                doors.openDoor0();
                            } catch (Exception ex) {
                                log.error("Ошибка: " + ex);
                            }
                        }
                    }
                } catch (Exception ex) {
                    log.error("Ошибка: " + ex);
                }
            });
            log.info("Подписка на топик Parking/MonitorDoor/# произошла успешно.");
        }catch(Exception ex){
            log.error("Ошибка: " + ex);
        }
    }

//    public static void startBackgroundMethods(){
//        new Thread(()->{ // проверка подключения к mqtt
//            while (true){
//                try {
//                    log.info("подключение к mqtt: " +(
//                            mqttService.getMqttClient().isConnected() ? "присутствует" : "отсутствует"));
//                    Thread.sleep(10000);
//                } catch (InterruptedException e) {
//                    throw new RuntimeException(e);
//                }
//            }
//        }).start();
}
