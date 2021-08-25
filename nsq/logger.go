package nsq

import (
	"time"

	"github.com/andrezhz/go-tools/log"
)

type PublishLogger interface {
	Println(format string, args ...interface{})
}

type DefaultNSQPublishLogger struct {
}

func (*DefaultNSQPublishLogger) Println(format string, args ...interface{}) {
	log.LOGFLOW(format, args...)
}

type ConsumerLogger interface {
	Output(calldepth int, s string) error
	Println(format string, args ...interface{})
}

type DefaultNSQConsumerLogger struct {
	ConsumerLogger
}

func (*DefaultNSQConsumerLogger) Println(format string, args ...interface{}) {
	log.LOGFLOW(format, args...)
}

func (DefaultNSQConsumerLogger) Output(calldepth int, s string) error {
	var err error
	log.PrintLog("nsq.consumer.log", time.Now(), &err, s, "")
	return nil
}
