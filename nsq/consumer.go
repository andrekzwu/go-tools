package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

var consumer *Consumer

type ConsumerEntry struct {
	Address string
	Logger  ConsumerLogger
}

type Consumer struct {
	address string
	lock    sync.RWMutex
	jobs    []*consumeJob
	logger  ConsumerLogger
	once    sync.Once
}

type Handler = nsq.Handler
type Message = nsq.Message
type HandlerFunc = nsq.HandlerFunc

type consumeJob struct {
	topic   string
	channel string
	handler Handler
	logger  ConsumerLogger
}

func (s *consumeJob) HandleMessage(nsqMsg *Message) error {
	t := time.Now()
	defer func() {
		if err := recover(); err != nil {
			if s.logger != nil {
				s.logger.Println("consumer.panic.topic(%s).channel(%s).messageId(%s).messageBody(%s).err(%v).latency(%v)", s.topic, s.channel, string(nsqMsg.ID[:]), string(nsqMsg.Body), fmt.Errorf("%v", err), time.Now().Sub(t))
			}
			return
		}
	}()
	err := s.handler.HandleMessage(nsqMsg)
	if err != nil && s.logger != nil {
		s.logger.Println("consumer.error.topic(%s).channel(%s).messageId(%s).messageBody(%s).err(%v).latency(%v)", s.topic, s.channel, string(nsqMsg.ID[:]), string(nsqMsg.Body), err, time.Now().Sub(t))
	}
	return err
}

func NewConsumer(entry *ConsumerEntry) *Consumer {
	if entry.Address == "" {
		panic("address is empty")
	}
	return &Consumer{
		address: entry.Address,
		jobs:    make([]*consumeJob, 0),
		logger:  entry.Logger,
	}
}

func RegisterNSQConsumer(entry *ConsumerEntry) {
	consumer = NewConsumer(entry)
}

func AddJob(topic string, channel string, handler Handler) {
	consumer.AddJob(topic, channel, handler)
}

func (s *Consumer) AddJob(topic string, channel string, handler Handler) *Consumer {
	s.lock.Lock()
	s.jobs = append(s.jobs, &consumeJob{
		topic:   topic,
		channel: channel,
		handler: handler,
		logger:  s.logger,
	})
	s.lock.Unlock()
	return s
}

func Start() {
	consumer.Start()
}

func (s *Consumer) Start() {
	if s.address == "" {
		panic("server is empty")
	}
	s.once.Do(func() {
		for _, job := range s.jobs {
			go s.runJob(job)
		}
	})
}

func (s *Consumer) runJob(cj *consumeJob) {
	defer func() {
		if err := recover(); err != nil {
			//3s后重新启动监听
			time.Sleep(time.Duration(3000) * time.Millisecond)
			s.runJob(cj)
			if s.logger != nil {
				s.logger.Println("consumer.retry.run.topic(%s).channel(%s).error(%v)", cj.topic, cj.channel, err)
			}
		}
	}()

	config := nsq.NewConfig()
	addr := s.address

	config.MaxInFlight = 5 // the MaxInFlight greater or equal to real nodes
	consumer, err := nsq.NewConsumer(cj.topic, cj.channel, config)
	if err != nil {
		panic("delivery failed to connect nsq:" + err.Error())
	}
	consumer.SetLogger(s.logger, nsq.LogLevelWarning)
	consumer.AddHandler(cj)
	err = consumer.ConnectToNSQLookupd(addr)

	if err != nil && s.logger != nil {
		s.logger.Println("delivery failed to connect nsq:%s", err.Error())
	}
}
