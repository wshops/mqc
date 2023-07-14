package mqc

import "time"

// QOS describes the quality of service of an mqtt publish
type QOS byte

const (
	// AtMostOnce means the broker will deliver at most once to every subscriber - this means message delivery is not guaranteed
	AtMostOnce QOS = iota
	// AtLeastOnce means the broker will deliver a message at least once to every subscriber
	AtLeastOnce
	// ExactlyOnce means the broker will deliver a message exactly once to every subscriber
	ExactlyOnce
)

type Options struct {
	serverURL            string        // URL for the broker (use tcp://)
	clientID             string        // Client ID to use when connecting to server
	connectRetry         bool          // How long to wait between connection attempts (defaults to 10s)
	connectRetryInterval time.Duration // use second
	connectTimeout       time.Duration // use second
	username             string
	password             string
	keepAlive            time.Duration // Keepalive period in seconds
}

var clientOption *Options

func NewOptions(serverURL string, clientID string, username string, password string) *Options {
	clientOption = &Options{
		serverURL: serverURL,
		clientID:  clientID,
		username:  username,
		password:  password,
	}
	return clientOption
}

func (o *Options) SetKeepAlive(keepAlive time.Duration) *Options {
	if o == nil {
		panic("Options == nil , please checkout")
	}
	o.keepAlive = keepAlive
	return o
}

func (o *Options) SetConnectTimeout(connectTimeout time.Duration) *Options {
	if o == nil {
		panic("Options == nil , please checkout")
	}
	o.connectTimeout = connectTimeout
	return o
}

func (o *Options) SetConnectRetry(connectRetry bool, connectRetryInterval time.Duration) *Options {
	if o == nil {
		panic("Options == nil , please checkout")
	}
	o.connectRetry = connectRetry
	if connectRetry {
		o.connectRetryInterval = connectRetryInterval
	}
	return o
}
