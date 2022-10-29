package redis

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"world_generator_processor_service/config"
	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	"github.com/go-redis/redis/v8"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/sirupsen/logrus"
)

type consumer struct {
	client      *redis.Client
	ch          string
	handler     HandlerFunc
	concurrency int
}

func NewConsumer(cfg *config.Config, ch string, handler HandlerFunc, concurrency int) (Consumer, error) {
	opts, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "parsing redis url error for consumer")
	}
	if username := cfg.Redis.Username; username != "" {
		opts.Username = username
	}
	if password := cfg.Redis.Password; password != "" {
		opts.Password = password
	}

	client := redis.NewClient(opts)
	if concurrency < 1 {
		concurrency = 1
	}

	return &consumer{
		client:      client,
		ch:          ch,
		handler:     handler,
		concurrency: concurrency,
	}, nil
}

func (c *consumer) Consume(ctx context.Context) {
	log := logrus.New().WithFields(logrus.Fields{
		"channel": c.ch,
		"fn":      "Consume",
	})
	ctxRedis, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(c.concurrency)
	defer wg.Wait()

	for i := 0; i < c.concurrency; i++ {
		go c.worker(ctxRedis, i, wg)
	}
	log.Info("consumer up and running...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctxRedis.Done():
		log.Info("terminating: context cancelled")
	case <-sigterm:
		log.Info("terminating: via signal")
	}

	cancel()
	wg.Wait()
	if err := c.client.Close(); err != nil {
		log.Error("close client error", err)
	}
}

func (c *consumer) worker(ctx context.Context, n int, wg *sync.WaitGroup) {
	log := logrus.New().WithFields(logrus.Fields{
		"channel":   c.ch,
		"fn":        "worker",
		"worker_no": n,
	})
	log.Info("worker was started...")
	defer log.Info("worker was stopped")
	defer wg.Done()

	s := c.client.Subscribe(ctx, c.ch)
	var errCount int
	for {
		log.Info("message received...")
		msg, err := s.ReceiveMessage(ctx)
		if err != nil {
			errCount++
			log.Error("receiving message error", err)
			if errCount > 10 {
				return
			}
		}
		if ctxErr := ctx.Err(); ctxErr != nil {
			log.Error("ctx error", ctxErr)
			return
		}

		var message Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			errCount++
			log.Error("unmarshal msg error", err)
			if errCount > 10 {
				return
			}
		}
		ctx = traceIDTools.SetTraceID(ctx, message.TraceID)
		if err := c.handler(ctx, message.Data); err != nil {
			errCount++
			b, err := json.Marshal(err)
			if err != nil {
				errCount++
				continue
			}
			log.Error("handle message error", string(b))
			if errCount > 10 {
				return
			}
		}
		log.Info("message processed...")
	}
}

func (c *consumer) GetClient() *redis.Client {
	return c.client
}

func (c *consumer) GetHandler() HandlerFunc {
	return c.handler
}
