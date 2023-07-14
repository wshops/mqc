package mqc

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"time"
)

type Mqc struct {
	client mqtt.Client
	option *mqtt.ClientOptions
	log    *zap.SugaredLogger
}

var instance *Mqc

func New(opt *Options, logger *zap.SugaredLogger) {
	if opt == nil {
		panic("opt == nil, please checkout")
	}
	MqcClientOption := mqtt.NewClientOptions()
	MqcClientOption.AddBroker(opt.serverURL).SetClientID(opt.clientID)
	MqcClientOption.SetUsername(opt.username).SetPassword(opt.password)
	if opt.keepAlive != 0*time.Second {
		MqcClientOption.SetKeepAlive(opt.keepAlive)
	}
	if opt.connectTimeout != 0*time.Second {
		MqcClientOption.SetConnectTimeout(opt.connectTimeout)
	}
	if opt.connectRetry {
		MqcClientOption.SetConnectRetry(opt.connectRetry).SetConnectRetryInterval(opt.connectRetryInterval)
	}
	instance = &Mqc{
		client: mqtt.NewClient(MqcClientOption),
		log:    logger,
		option: MqcClientOption,
	}
	if token := instance.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
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

func Publish(topic string, qos QOS, retain bool, payload []byte) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	return instance.publish(topic, qos, retain, payload)
}

func RegisterSubscriber(topic string, qos QOS, handler mqtt.MessageHandler) error {
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
