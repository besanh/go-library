package redismq

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

type Rcfg struct {
	Address  string
	Username string
	Password string
	DB       int
}

var RMQ *RMQConnection

const (
	tag           = "rmq"
	pollDuration  = 100 * time.Millisecond
	prefetchLimit = 1000
)

type RMQConnection struct {
	RedisClient *redis.Client
	Config      Rcfg
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

func NewRMQ(config Rcfg) *RMQConnection {
	poolSize := runtime.NumCPU() * 4
	errChan := make(chan error, 10)
	go logErrors(errChan)
	client := redis.NewClient(&redis.Options{
		Addr:            config.Address,
		Password:        config.Password,
		DB:              config.DB,
		PoolSize:        poolSize,
		PoolTimeout:     time.Duration(20) * time.Second,
		ReadTimeout:     time.Duration(20) * time.Second,
		WriteTimeout:    time.Duration(20) * time.Second,
		ConnMaxIdleTime: time.Duration(20) * time.Second,
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
	}
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
