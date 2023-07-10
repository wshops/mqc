package mqc

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
)

func (m *Mqc) Publish(topic string, payload []byte, qos QOS, retain ...bool) error {
	retainFlag := false
	if len(retain) > 0 {
		retainFlag = retain[0]
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pb, err := m.cm.Publish(ctx, &paho.Publish{
		QoS:     byte(qos),
		Retain:  retainFlag,
		Topic:   topic,
		Payload: payload,
	})
	if err != nil {
		return err
	}
	if err != nil {
		m.options.logger.Error("MQ publish error: ", err)
	} else if pb.ReasonCode != 0 && pb.ReasonCode != 16 { // 16 = Server received message but there are no subscribers
		m.options.logger.Error("MQ publish error: ", pb.ReasonCode)
	}
	m.options.logger.Debug(fmt.Sprintf("MQ message published: %s", topic))
	return nil
}
