package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wshops/mqc"
	"github.com/wshops/zlog"
	"time"
)

func main() {
	zlog.New(zlog.LevelDev)

	options := mqc.NewMqcOptions().AddBroker("tcp://100.100.10.13:1883").SetClientID("go_test_client")
	options.SetUsername("goTest").SetPassword("123456")
	options.SetKeepAlive(30 * time.Second)
	options.SetPingTimeout(1 * time.Second)

	mqc.New(options, zlog.Log())
	mqc.ServerConnect()
	err := mqc.RegisterSubscriber("test1", mqc.ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", message.Topic())
		fmt.Printf("MSG: %s\n", message.Payload())
	})
	if err != nil {
		panic(err)
	}
	mqc.ServerConnect()
	err = mqc.Publish("test1", mqc.ExactlyOnce, false, []byte("this is a test message!"))
	if err != nil {
		panic(err)
	}
	defer mqc.ServerDisconnect(0)
	time.Sleep(time.Second * 5)
}
