package mqc

import "github.com/eclipse/paho.golang/paho"

type Message struct {
	msgPackage *paho.Publish
}

type MsgHandler func(Message) error

func (m *Message) Topic() string {
	return m.msgPackage.Topic
}

func (m *Message) Payload() []byte {
	return m.msgPackage.Payload
}

func (m *Message) Qos() QOS {
	return QOS(m.msgPackage.QoS)
}

func (m *Message) IsDuplicate() bool {
	return m.msgPackage.Packet().Duplicate
}

func (m *Message) IsRetain() bool {
	return m.msgPackage.Packet().Retain
}

func (m *Message) PacketID() uint16 {
	return m.msgPackage.Packet().PacketID
}

func (m *Mqc) RegisterSubscriber(topic string, qos QOS, handler MsgHandler) {
	m.subscribeOptions[topic] = paho.SubscribeOptions{
		QoS: byte(qos),
	}
	m.router.RegisterHandler(topic, func(publish *paho.Publish) {
		err := handler(Message{msgPackage: publish})
		if err != nil {
			m.options.logger.Error("MQ message handler error, err: ", err)
			return
		}
	})
}
