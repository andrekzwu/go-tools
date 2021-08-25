package main

import (
	"context"
	"fmt"

	"github.com/andrezhz/go-tools/nsq"
)

func main() {
	// setup nsq consumer
	SetupNSQ()
	ctx := context.Background()
	<-ctx.Done()
}

func SetupNSQ() {
	nsq.RegisterNSQConsumer(&nsq.ConsumerEntry{
		Address: "xxxxx:4161",
	})
	// use function
	nsq.AddJob("GO_TOOLS_TEST", "CHANNEL_ONE", nsq.HandlerFunc(HandleMessageOne))
	// use struct
	nsq.AddJob("GO_TOOLS_TEST", "CHANNEL_TWO", new(TestNSQConsumer))
	nsq.AddJob("GO_TOOLS_TEST", "CHANNEL_TWO", new(TestNSQConsumer))
	// start
	nsq.Start()
}

func HandleMessageOne(message *nsq.Message) error {
	fmt.Println(string(message.ID[:]), string(message.Body))
	return nil
}

type TestNSQConsumer struct {
}

func (*TestNSQConsumer) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.ID[:]), string(message.Body))
	//panic("test error")
	return nil
}
