package mqc

import mqtt "github.com/eclipse/paho.mqtt.golang"

// QOS describes the quality of service of an mqtt publish
type QOS byte

const (
	// AtMostOnce means the broker will deliver at most once to every subscriber - this means message delivery is not guaranteed
	AtMostOnce byte = iota
	// AtLeastOnce means the broker will deliver a message at least once to every subscriber
	AtLeastOnce
	// ExactlyOnce means the broker will deliver a message exactly once to every subscriber
	ExactlyOnce
)

func NewMqcOptions() *mqtt.ClientOptions {
	return mqtt.NewClientOptions()
}

//
//type MqcOptions struct {
//	serverURL            string        // URL for the broker (schemes supported include 'mqtt' and 'tls')
//	clientID             string        // Client ID to use when connecting to server
//	keepAlive            time.Duration // Keepalive period in seconds (the maximum time interval that is permitted to elapse between the point at which the Client finishes transmitting one MQTT Control Packet and the point it starts sending the next)
//	connectRetry         bool          /// How long to wait between connection attempts (defaults to 10s)
//	connectRetryInterval time.Duration
//	connectTimeout       time.Duration
//	authUsername         string
//	authPassword         string
//}
//
//var opt MqcOptions
//
//func NewMqcOptions(serverURL string, clientID string, keepAlive time.Duration, authUsername string, authPassword string) MqcOptions {
//	opt = MqcOptions{
//		serverURL:    serverURL,
//		clientID:     clientID,
//		keepAlive:    keepAlive,
//		authUsername: authUsername,
//		authPassword: authPassword,
//	}
//	return opt
//}
//
//func (o *MqcOptions) SetConnectTimeout(connectTimeout time.Duration) *MqcOptions {
//	o.connectTimeout = connectTimeout
//	return o
//}
//
//func (o *MqcOptions) SetConnectRetry(connectRetry bool, connectRetryInterval time.Duration) *MqcOptions {
//	o.connectRetry = connectRetry
//	if connectRetry {
//		o.connectRetryInterval = connectRetryInterval
//	}
//	return o
//}
