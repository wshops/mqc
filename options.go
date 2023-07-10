package mqc

import (
	"github.com/wshops/zlog"
	"go.uber.org/zap"
	"net/url"
	"time"
)

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

type MqcOptions struct {
	logger            *zap.SugaredLogger
	serverURL         *url.URL      // URL for the broker (schemes supported include 'mqtt' and 'tls')
	clientID          string        // Client ID to use when connecting to server
	keepAlive         uint16        // Keepalive period in seconds (the maximum time interval that is permitted to elapse between the point at which the Client finishes transmitting one MQTT Control Packet and the point it starts sending the next)
	connectRetryDelay time.Duration /// How long to wait between connection attempts (defaults to 10s)
	connectTimeout    time.Duration // How long to wait for the connection process to complete (defaults to 10s)
	authUsername      string
	authPassword      string
}

type Option func(options *MqcOptions)

var defaultOption = &MqcOptions{
	logger:            zlog.New(zlog.LevelProd).SugaredLogger,
	serverURL:         &url.URL{Scheme: "mqtt", Host: "localhost:1883"},
	clientID:          "wshops-mqtt-client",
	keepAlive:         10,
	connectRetryDelay: 10 * time.Second,
	connectTimeout:    10 * time.Second,
	authUsername:      "",
	authPassword:      "",
}

func WithLogger(logger *zap.SugaredLogger) Option {
	return func(o *MqcOptions) {
		o.logger = logger
	}
}

func WithServerURL(urlStr string) Option {
	u, err := url.Parse(urlStr)
	return func(o *MqcOptions) {
		if err != nil {
			o.serverURL = defaultOption.serverURL
		} else {
			o.serverURL = u
		}
	}
}

func WithClientID(clientID string) Option {
	return func(o *MqcOptions) {
		o.clientID = clientID
	}
}

func WithKeepAlive(keepAlive uint16) Option {
	return func(o *MqcOptions) {
		o.keepAlive = keepAlive
	}
}

func WithConnectRetryDelay(connectRetryDelay time.Duration) Option {
	return func(o *MqcOptions) {
		o.connectRetryDelay = connectRetryDelay
	}
}

func WithConnectTimeout(connectTimeout time.Duration) Option {
	return func(o *MqcOptions) {
		o.connectTimeout = connectTimeout
	}
}

func WithAuth(username string, password string) Option {
	return func(o *MqcOptions) {
		o.authUsername = username
		o.authPassword = password
	}
}
