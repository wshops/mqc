package mqc

import (
	"context"
	"github.com/eclipse/paho.golang/paho"
	"time"
)

type router struct {
}

func (mqc *Mqc) RegisterSubscriber(topic string, qos byte) error {
	instance.subscribeOptions[topic] = paho.SubscribeOptions{
		QoS: qos,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := mqc.connectionManager.Subscribe(ctx, &paho.Subscribe{
		Subscriptions: instance.subscribeOptions,
	})
	if err != nil {
		return err
	}

	return nil
}
