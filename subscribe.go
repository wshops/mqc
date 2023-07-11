package mqc

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (m *Mqc) registerSubscriber(topic string, qos byte, handler mqtt.MessageHandler) error {
	if !m.client.IsConnected() {
		ServerConnect()
	}

	token := m.client.Subscribe(topic, qos, handler)

	if token.Wait() && token.Error() != nil {
		m.log.Error(token.Error())
		return token.Error()
	}

	m.log.Debug(fmt.Sprintf("MQ message subscribed, topic is %s", topic))
	return nil
}

func (m *Mqc) unsubscribe(topic string) error {
	if !m.client.IsConnected() {
		ServerConnect()
	}

	token := m.client.Unsubscribe(topic)

	if token.Wait() && token.Error() != nil {
		m.log.Error(token.Error())
		return token.Error()
	}

	m.log.Debug(fmt.Sprintf("MQ message has not subscribed, topic is %s", topic))
	return nil
}
