package nsq

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
)

var producer *Producer

type ProducerEntry struct {
	Addr   string
	Logger PublishLogger
}

type Producer struct {
	producer *nsq.Producer
	logger   PublishLogger
}

func NewProducer(entry *ProducerEntry) *Producer {
	if err := checkNSQEntry(entry); err != nil {
		panic("invalid host")
	}
	_producer, err := nsq.NewProducer(entry.Addr, nsq.NewConfig())
	//  try to ping
	err = _producer.Ping()
	if nil != err {
		_producer.Stop()
		msg := fmt.Sprint("failed to connect nsq,addr =", entry.Addr, "err=", err.Error())
		panic(msg)
	}
	return &Producer{producer: _producer, logger: entry.Logger}
}

func RegisterNSQProducer(entry *ProducerEntry) {
	producer = NewProducer(entry)
}

func Publish(topic string, message []byte) error {
	return producer.Publish(topic, message)
}

func (producer *Producer) Publish(topic string, message []byte) error {
	if producer == nil {
		return errors.New("producer is nil!")
	}
	if message == nil || len(message) == 0 {
		return errors.New("The message is empty")
	}
	var err error
	err = producer.producer.Publish(topic, message)
	if err != nil && producer.logger != nil {
		producer.logger.Println("producer.published.topic(%s).message(%s).error(%v)", topic, string(message), err)
	}
	return err
}

func MultiPublish(topic string, message [][]byte) error {
	return producer.MultiPublish(topic, message)
}

func (producer *Producer) MultiPublish(topic string, message [][]byte) error {
	if producer == nil {
		return errors.New("producer is nil!")
	}
	if message == nil || len(message) == 0 {
		return errors.New("The message is empty")
	}
	var err error
	err = producer.producer.MultiPublish(topic, message)
	if err != nil && producer.logger != nil {
		producer.logger.Println("producer.published.topic(%s).message(%s).error(%v)", topic, transMessage2Str(message), err)
	}
	return err
}

func transMessage2Str(message [][]byte) string {
	var buff bytes.Buffer
	for i, item := range message {
		if i != 0 {
			buff.WriteString(",")
		}
		buff.Write(item)
	}
	return buff.String()
}

func checkNSQEntry(entry *ProducerEntry) error {
	if entry.Addr == "" {
		return errors.New("invalid addr")
	}
	return nil
}
