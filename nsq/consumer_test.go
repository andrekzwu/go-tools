package nsq

import (
	"context"
	"fmt"
	"testing"

	"github.com/nsqio/go-nsq"
)

func init() {
	RegisterNSQConsumer(&ConsumerEntry{
		Address: "xxxxxxxxx:4161",
	})
}

func TestAddJob(t *testing.T) {
	ctx := context.Background()
	// add job
	AddJob("GO_TOOLS_TEST", "CHANNEL_ONE", nsq.HandlerFunc(HandleMessage))
	// start
	Start()
	<-ctx.Done()
}

func HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.ID[:]), string(message.Body))
	return nil
}
