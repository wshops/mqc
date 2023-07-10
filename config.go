package mqc

import (
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"net/url"
)

const (
	// AtMostOnce means the broker will deliver at most once to every subscriber - this means message delivery is not guaranteed
	AtMostOnce = iota
	// AtLeastOnce means the broker will deliver a message at least once to every subscriber
	AtLeastOnce
	// ExactlyOnce means the broker will deliver a message exactly once to every subscriber
	ExactlyOnce
)

type MqcConfig struct {
	Cfg    autopaho.ClientConfig
	router paho.SingleHandlerRouter
}

func NewMqcConfig() *MqcConfig {
	return &MqcConfig{}
}

func (m *MqcConfig) InitMqttConfig(brokerUrl string, keepAlive uint16, clientID string) *MqcConfig {
	u, err := url.Parse(brokerUrl)
	if err != nil {
		panic(err)
	}
	m.Cfg.BrokerUrls = []*url.URL{u}
	m.Cfg.KeepAlive = keepAlive
	m.Cfg.ClientConfig = paho.ClientConfig{
		ClientID: clientID,
		Router: paho.NewSingleHandlerRouter(func(publish *paho.Publish) {
			fmt.Print(publish)
		}),
	}

	return m
}

//func (m *MqcConfig) RegisterMessageHandler(topic string, handler paho.MessageHandler) *MqcConfig {
//	m.Router.RegisterHandler("topic", handler)
//
//	return m
//}
