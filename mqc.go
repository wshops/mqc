package mqc

import (
	"context"
	"errors"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"net/url"
	"time"
)

type Mqc struct {
	options          *MqcOptions
	subscribeOptions map[string]paho.SubscribeOptions
	router           *paho.StandardRouter
	cm               *autopaho.ConnectionManager
	clientConfig     *autopaho.ClientConfig
	cancelConnection context.CancelFunc
}

var instance *Mqc

func New(opt ...Option) *Mqc {
	if instance != nil {
		return instance
	}
	options := defaultOption
	for _, o := range opt {
		o(options)
	}
	m := &Mqc{
		options:          options,
		subscribeOptions: make(map[string]paho.SubscribeOptions),
		router:           paho.NewStandardRouter(),
	}
	pahoClientConfig := &autopaho.ClientConfig{
		BrokerUrls: []*url.URL{
			options.serverURL,
		},
		KeepAlive:         options.keepAlive,
		ConnectRetryDelay: options.connectRetryDelay,
		ConnectTimeout:    options.connectTimeout,
		OnConnectionUp:    m.connectionUpProcessor,
		OnConnectError:    m.connectionErrorProcessor,
		ClientConfig: paho.ClientConfig{
			ClientID:           options.clientID,
			PingHandler:        m.pingProcessor(),
			Router:             m.router,
			PacketTimeout:      10 * time.Second,
			OnServerDisconnect: m.serverDisconnectProcessor,
			OnClientError: func(err error) {
				options.logger.Error("MQ client error: ", err)
			},
		},
	}
	m.clientConfig = pahoClientConfig
	instance = m
	return m
}

func (m *Mqc) pingProcessor() *paho.PingHandler {
	return paho.DefaultPingerWithCustomFailHandler(func(err error) {
		m.options.logger.Error("MQ ping failed")
	})
}

func (m *Mqc) serverDisconnectProcessor(disconnect *paho.Disconnect) {
	if disconnect.Properties != nil {
		m.options.logger.Error("MQ server requested disconnect; reason code: ", disconnect.ReasonCode, ", reason: ", disconnect.Properties.ReasonString)
	} else {
		m.options.logger.Error("MQ server requested disconnect; reason code: ", disconnect.ReasonCode)
	}
}

func (m *Mqc) connectionUpProcessor(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
	m.options.logger.Info("MQ connection up")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if _, err := cm.Subscribe(ctx, &paho.Subscribe{
		Subscriptions: m.subscribeOptions,
	}); err != nil {
		m.options.logger.Error("MQ subscribe failed, err: ", err)
		return
	}
	m.options.logger.Info("MQ subscribe success")
}

func Start() {
	if instance == nil {
		return
	}
	instance.Start()
}

func Stop() {
	if instance == nil {
		return
	}
	instance.Stop()
}

func RegisterSubscriber(topic string, qos QOS, handler MsgHandler) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	instance.RegisterSubscriber(topic, qos, handler)
	return nil
}

func Publish(topic string, qos QOS, payload []byte, retain ...bool) error {
	if instance == nil {
		return errors.New("MQC not initialized")
	}
	return instance.Publish(topic, payload, qos, retain...)
}

func (m *Mqc) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancelConnection = cancel
	cm, err := autopaho.NewConnection(ctx, *m.clientConfig)
	if err != nil {
		m.options.logger.Fatal("MQ connection failed, err: ", err)
	}
	m.cm = cm
}

func (m *Mqc) Stop() {
	for k, _ := range m.subscribeOptions {
		m.router.UnregisterHandler(k)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := m.cm.Disconnect(ctx)
	if err != nil {
		m.options.logger.Error("MQ disconnect failed, err: ", err)
	}
	m.cancelConnection()
}

func (m *Mqc) connectionErrorProcessor(err error) {
	m.options.logger.Error(err)
}
