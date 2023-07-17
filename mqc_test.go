package mqc

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
	"github.com/wshops/zlog"
	"strconv"
	"testing"
)

func BenchmarkParallel(b *testing.B) {
	zlog.New(zlog.LevelDev)

	o := NewOptions("tcp://100.100.10.13:1883", "go_test_client", "goTest", "123456")
	New(o, zlog.Log())
	RegisterSubscriber("test1", ExactlyOnce, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MSG: %s\n", message.Payload())
	})
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; i < b.N; i++ {
			fmt.Print(i)
			Publish("test1", ExactlyOnce, false, []byte(strconv.FormatInt(int64(i), 10)))
		}
	})
}

func TestRegisterMultipleSubscriber(t *testing.T) {
	zlog.New()
	o := NewOptions("tcp://100.100.10.13:1883", "go_test_client", "goTest", "123456")
	New(o, zlog.Log())
	topics := map[string]byte{
		"test1": byte(ExactlyOnce),
		"test2": byte(ExactlyOnce),
	}

	var receiveMessage []string

	RegisterMultipleSubscriber(topics, func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MSG: %s\n", message.Payload())
		receiveMessage = append(receiveMessage, string(message.Payload()))
	})

	Publish("test1", ExactlyOnce, false, []byte("test1:123"))
	Publish("test2", ExactlyOnce, false, []byte("test2:456"))
	fmt.Println(receiveMessage)
	assert.Equal(t, "test1:123", receiveMessage[0])
	assert.Equal(t, "test2:456", receiveMessage[1])
}
