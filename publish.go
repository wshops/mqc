package mqc

import (
	"fmt"
)

func (m *Mqc) publish(topic string, qos QOS, retain bool, payload []byte) error {
	if !m.client.IsConnected() {
		ServerConnect()
	}
	token := m.client.Publish(topic, byte(qos), retain, payload)
	token.Wait()
	if token.Error() != nil {
		m.log.Error(token.Error())
		return token.Error()
	}

	m.log.Debug(fmt.Sprintf("MQ message published: %s", topic))
	return nil
}
