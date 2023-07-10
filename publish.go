package mqc

import (
	"context"
	"github.com/eclipse/paho.golang/paho"
	"time"
)

func (mqc *Mqc) PublishMessage(topic string, qos byte, msg []byte) error {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pr, err := instance.connectionManager.Publish(ctx, &paho.Publish{
		QoS:     qos,
		Topic:   topic,
		Payload: msg,
	})
	if err != nil {
		return err
	}

	if pr.ReasonCode != 0 && pr.ReasonCode != 16 { // 16 // = Server received message but there are no subscribers
		instance.log.Warnf("reason code %d received\n", pr.ReasonCode)
	}

	return nil
}
