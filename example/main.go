package main

import (
	"fmt"
	"github.com/wshops/mqc"
	"github.com/wshops/zlog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	zlog.New(zlog.LevelDev)
	mqc.New(
		mqc.WithServerURL("mqtt://173.82.219.121:1883"),
		mqc.WithClientID("wshops-mqtt-client-test"),
		mqc.WithAuth("anxuanzi", "anxuanzi"),
		mqc.WithKeepAlive(10),
		mqc.WithConnectRetryDelay(10*time.Second),
		mqc.WithLogger(zlog.Log()),
	)

	err := mqc.RegisterSubscriber("aaa/bbb/ccc", mqc.AtMostOnce, func(message mqc.Message) error {
		fmt.Println("receive message:", message.Payload())
		return nil
	})
	if err != nil {
		return
	}

	mqc.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)
	<-sig
	fmt.Println("signal caught - exiting")
	mqc.Stop()
	fmt.Println("shutdown complete")
}
