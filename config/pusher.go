package config

import (
	"fmt"
	pusher "github.com/pusher/pusher-http-go/v5"
	"os"
)

func Pusher(data interface{}, prefix string) {
	channel, _ := os.LookupEnv("CHANNEL_NAME")
	event, _ := os.LookupEnv("EVENT_NAME")
	pusherId, _ := os.LookupEnv("PUSHER_APP_ID")
	pusherKey, _ := os.LookupEnv("PUSHER_APP_KEY")
	pusherSecret, _ := os.LookupEnv("PUSHER_APP_SECRET")
	pusherCluster, _ := os.LookupEnv("PUSHER_APP_CLUSTER")
	Client := pusher.Client{
		AppID:   pusherId,
		Key:     pusherKey,
		Secret:  pusherSecret,
		Cluster: pusherCluster,
		Secure:  true,
	}
	err := Client.Trigger(channel, event+"_"+prefix, data)
	if err != nil {
		fmt.Println("pusher log :", err.Error())
	}
}
