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
	o := mqc.NewOptions("tcp://100.100.10.13:1883", "go_test_client", "goTest", "123456")
	mqc.New(o, zlog.Log())
	mqc.RegisterSubscriber("test1", mqc.ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", message.Topic())
		fmt.Printf("MSG: %s\n", message.Payload())
	})
	for i := 0; i < 100; i++ {
		mqc.Publish("test1", mqc.ExactlyOnce, false, []byte("this is a test message!"))
	}
	defer mqc.ServerDisconnect(0)
	time.Sleep(time.Second * 5)
}
