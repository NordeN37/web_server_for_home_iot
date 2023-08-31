package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
)

const (
	TopicRoom = "Room"
)

type IMQTT interface {
	MQTTPublish(topic string, msg []byte)
	MQTTSubscribe(topic string)
	Close()
}

type mqttClient struct {
	MQTT.Client
}

func NewClient(endpoint, clientId, username, password string) (IMQTT, error) {
	MQTT.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	MQTT.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	MQTT.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	MQTT.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(endpoint) // Update with your Mosquitto broker address
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &mqttClient{client}, nil
}

func (m *mqttClient) MQTTPublish(topic string, msg []byte) {
	t := m.Publish(topic, 0, false, msg)
	go func() {
		_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
		if t.Error() != nil {
			fmt.Println(t.Error())
		}
	}()
}

func (m *mqttClient) Close() {
	m.Disconnect(250)
	m.Close()
}

func (m *mqttClient) MQTTSubscribe(topic string) {
	m.Subscribe(topic, 0, func(c MQTT.Client, m MQTT.Message) {
		log.Println("Subscribe: ", string(m.Payload()))
		m.Ack()
	})
}
