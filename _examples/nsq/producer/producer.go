package main

import (
	"fmt"
	"time"

	"github.com/andrezhz/go-tools/nsq"
)

func main() {
	// setup nsq producer
	SetupNSQ()
	// publish test msg
	Publish()
}

func Publish() {
	for i := 0; i < 10; i++ {
		nsq.Publish("GO_TOOLS_TEST", []byte(fmt.Sprintf("Hello World!%d", i)))
		time.Sleep(time.Second * 1)
	}
	fmt.Println("Done")
}

func SetupNSQ() {
	// init nsq producer
	nsq.RegisterNSQProducer(&nsq.ProducerEntry{
		Addr: "xxxxx:4150",
	})
}
