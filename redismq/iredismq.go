package redismq

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adjust/rmq/v5"
)

type IRedisMQ interface {
	AddQueue(name string, handler rmq.ConsumerFunc, numConsumers int) error
	RemoveQueue(name string) error
	IsHasQueue(name string) bool
}

var (
	ErrQueueIsExist    error = errors.New("queue is existed")
	ErrQueueIsNotExist error = errors.New("queue is not existed")

	RedisMQ IRedisMQ
)

func (srv *RMQServer) AddQueue(name string, handler rmq.ConsumerFunc, numConsumers int) error {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if _, ok := srv.Queues[name]; ok {
		return ErrQueueIsExist
	}
	queue, err := srv.conn.OpenQueue(name)
	if err != nil {
		return err
	}
	srv.Queues[name] = queue
	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		return err
	}
	for range numConsumers {
		if _, err := queue.AddConsumerFunc(tag, handler); err != nil {
			return err
		}
	}
	return nil
}

type Consumer struct {
}

func (c *Consumer) Consume(delivery rmq.Delivery) {
	log.Println("Received message: ", delivery.Payload())
	delivery.Ack()
}

func (conn *RMQConnection) Close() {
	<-conn.Conn.StopAllConsuming()
	cleaner := rmq.NewCleaner(conn.Conn)
	i, err := cleaner.Clean()
	if err != nil {
		log.Println(err)
	}
	log.Println("Cleaned", i, "messages")
}

func (srv *RMQServer) Run() {
	srv.waitForSignals()
}

func (srv *RMQServer) waitForSignals() {
	log.Println("Waiting for signals...")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()

	<-srv.conn.StopAllConsuming()
}

func (srv *RMQServer) RemoveQueue(name string) error {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if _, ok := srv.Queues[name]; !ok {
		return ErrQueueIsNotExist
	}
	q, err := srv.conn.OpenQueue(name)
	if err != nil {
		return err
	}
	q.Destroy()
	srv.Queues[name].StopConsuming()
	// srv.Queues[name].PurgeReady()
	// srv.Queues[name].PurgeRejected()
	delete(srv.Queues, name)
	return nil
}

func (srv *RMQServer) IsHasQueue(name string) bool {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	_, ok := srv.Queues[name]
	return ok
}

func (c *RMQClient) PublishBytes(queueName string, payload []byte) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	queue, ok := c.Queues[queueName]
	if !ok {
		queue, err = c.conn.OpenQueue(queueName)
		if err != nil {
			return err
		}
		c.Queues[queueName] = queue
		fmt.Println("new queue: ", queueName)
	}
	return queue.PublishBytes(payload)
}

func (c *RMQClient) Publish(queueName string, payload string) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	queue, ok := c.Queues[queueName]
	if !ok {
		queue, err = c.conn.OpenQueue(queueName)
		if err != nil {
			return err
		}
		c.Queues[queueName] = queue
		fmt.Println("new queue: ", queueName)
	}
	return queue.Publish(payload)
}
