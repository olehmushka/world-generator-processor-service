package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"world_generator_processor_service/config"
	traceIDTools "world_generator_processor_service/core/tools/trace_id"

	"github.com/go-redis/redis/v8"
	timeTools "github.com/olehmushka/golang-toolkit/time_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.uber.org/fx"
)

type publisher struct {
	client *redis.Client
}

func NewPublisher(cfg *config.Config) (Publisher, error) {
	opts, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "parsing redis url error for publisher")
	}
	if username := cfg.Redis.Username; username != "" {
		opts.Username = username
	}
	if password := cfg.Redis.Password; password != "" {
		opts.Password = password
	}

	client := redis.NewClient(opts)
	return &publisher{
		client: client,
	}, nil
}

var PublisherModule = fx.Options(
	fx.Provide(NewPublisher),
)

func (p *publisher) Publish(ctx context.Context, ch string, v []byte) error {
	msg := Message{
		Data:      v,
		Timestamp: timeTools.TimeToString(time.Now()),
		TraceID:   traceIDTools.GetTraceID(ctx),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return wrapped_error.NewInternalServerError(err, "can not marshal redis message")
	}

	if err := p.client.Publish(ctx, ch, b).Err(); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not publish redis msg (channel=%s, msg=%+v)", ch, msg))
	}

	return nil
}
