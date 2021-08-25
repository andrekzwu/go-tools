package nsq

import (
	"fmt"
	"testing"
)

func init() {
	RegisterNSQProducer(&NSQProducerEntry{
		Addr: "xxxxxxxxxx:4150",
	})
}

func TestPublish(t *testing.T) {
	err := Publish("GO_TOOLS_TEST", []byte(fmt.Sprintf("Hello World!")))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiPublish(t *testing.T) {
	bytes := make([][]byte, 0, 2)
	bytes = append(bytes, []byte(fmt.Sprintf("Hello World! 1")))
	bytes = append(bytes, []byte(fmt.Sprintf("Hello World! 2")))
	err := MultiPublish("GO_TOOLS_TEST", bytes)
	if err != nil {
		t.Fatal(err)
	}
}
