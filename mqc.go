package mqc

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Mqc struct {
	client mqtt.Client
	log    *zap.SugaredLogger
}

var instance *Mqc

func New(opt *mqtt.ClientOptions, logger *zap.SugaredLogger) {
	instance = &Mqc{
		client: mqtt.NewClient(opt),
		log:    logger,
	}
	if token := instance.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	instance.client.Connect()
}

// ServerDisconnect Disconnect will end the connection with the server,
// but not before waiting the specified number of milliseconds
// to wait for existing work to be completed.
func ServerDisconnect(quiesce uint) {
	instance.client.Disconnect(quiesce)
}

// ServerConnect will create a connection to the message broker, by default
func ServerConnect() {
	if token := instance.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func Publish(topic string, qos byte, retain bool, payload []byte) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	return instance.publish(topic, qos, retain, payload)
}

func RegisterSubscriber(topic string, qos byte, handler mqtt.MessageHandler) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	return instance.registerSubscriber(topic, qos, handler)
}

func Unsubscribe(topic string) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	return instance.unsubscribe(topic)
}
