package pubsub

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/tranvannghia021/gocore/vars"
	"log"
)

type PubSub interface {
	Send(payload interface{}) error
	Listen() chan string
}
type pubsub struct {
	redis   *redis.Client
	channel string
}

func NewPubSub(channel string) PubSub {
	return &pubsub{
		redis:   vars.Redis,
		channel: channel,
	}
}
func (p *pubsub) Send(payload interface{}) error {
	payload, _ = json.Marshal(payload)
	return p.redis.Publish(p.channel, payload).Err()
}

func (p *pubsub) Listen() chan string {
	var result = make(chan string)
	subscriber := p.redis.Subscribe(p.channel)
	go func() {
		for {
			msg, err := subscriber.ReceiveMessage()
			if err != nil {
				log.Fatal(err)
			}
			result <- msg.String()
		}
	}()
	return result
}
