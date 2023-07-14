package mqc

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wshops/zlog"
	"strconv"
	"testing"
)

func BenchmarkParallel(b *testing.B) {
	zlog.New(zlog.LevelDev)

	o := NewOptions("tcp://100.100.10.13:1883", "go_test_client", "goTest", "123456")
	New(o, zlog.Log())
	err := RegisterSubscriber("test1", ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MSG: %s\n", message.Payload())
	})
	if err != nil {
		panic(err)
	}
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < b.N; i++ {
			fmt.Print(i)
			err = Publish("test1", ExactlyOnce, false, []byte(strconv.FormatInt(int64(i), 10)))
			if err != nil {
				panic(err)
			}
		}
	})
}
