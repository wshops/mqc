package mqc

import (
	"context"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"go.uber.org/zap"
	"time"
)

type Mqc struct {
	log               *zap.SugaredLogger
	connectionManager *autopaho.ConnectionManager
	subscribeOptions  map[string]paho.SubscribeOptions
}

var instance *Mqc

func New(logger *zap.SugaredLogger, mqcConfig *MqcConfig) *Mqc {
	if instance != nil {
		instance.log.Warn("mqc instance already exists")
		return instance
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cm, err := autopaho.NewConnection(ctx, mqcConfig.Cfg)

	if err != nil {
		panic(err)
	}

	instance = &Mqc{
		log:               logger,
		connectionManager: cm,
		subscribeOptions:  make(map[string]paho.SubscribeOptions),
	}

	return instance
}
