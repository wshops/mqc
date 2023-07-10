package mqc

import (
	"github.com/wshops/zlog"
	"testing"
)

func TestReceiveMessage(t *testing.T) {
	mqcConfig := NewMqcConfig()
	mqcConfig.InitMqttConfig("tcp://100.100.10.13:1883", 30, "mqc_test_publisher")
	//mqcConfig.RegisterMessageHandler("test", func(publish *paho.Publish) {
	//	fmt.Print(publish.Packet().String())
	//	fmt.Print(publish.String())
	//})
	mqcConfig.Cfg.SetUsernamePassword("goTest", []byte("123456"))
	zlog.New(zlog.LevelDev)
	mqc := New(zlog.Log(), mqcConfig)
	err := mqc.RegisterSubscriber("test", 2)
	if err != nil {
		panic(err)
	}
	err = mqc.PublishMessage("test", 2, []byte("this is test message"))
	if err != nil {
		panic(err)
	}
}
