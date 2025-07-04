package redismq

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address         string
	Username        string
	Password        string
	DB              int
	PoolSize        int
	PoolTimeout     time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ConnMaxIdleTime time.Duration
}

var RMQ *RMQConnection

const (
	tag           = "rmq"
	pollDuration  = 100 * time.Millisecond
	prefetchLimit = 1000
)

type RMQConnection struct {
	RedisClient *redis.Client
	Config      Config
	Conn        rmq.Connection
	Queues      map[string]rmq.Queue
	Server      *RMQServer
	Client      *RMQClient
}

type RMQServer struct {
	mu     sync.Mutex
	conn   rmq.Connection
	Queues map[string]rmq.Queue
}

type RMQClient struct {
	mu     sync.Mutex // protects Queues
	conn   rmq.Connection
	Queues map[string]rmq.Queue
}

func NewRMQ(config Config) (*RMQConnection, error) {
	poolSize := runtime.NumCPU() * 4
	errChan := make(chan error, 10)
	go logErrors(errChan)
	client := redis.NewClient(&redis.Options{
		Addr:            config.Address,
		Username:        config.Username,
		Password:        config.Password,
		DB:              config.DB,
		PoolSize:        poolSize,
		PoolTimeout:     config.PoolTimeout,
		DialTimeout:     config.PoolTimeout,
		ReadTimeout:     config.ReadTimeout,
		WriteTimeout:    config.WriteTimeout,
		ConnMaxIdleTime: config.ConnMaxIdleTime,
	})
	connection, err := rmq.OpenConnectionWithRedisClient(tag, client, errChan)
	if err != nil {
		log.Fatal(err)
	}

	return &RMQConnection{
		RedisClient: client,
		Config:      config,
		Conn:        connection,
		Server: &RMQServer{
			Queues: make(map[string]rmq.Queue),
			conn:   connection,
		},
		Client: &RMQClient{
			Queues: make(map[string]rmq.Queue),
			conn:   connection,
		},
	}, nil
}
func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}
