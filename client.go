package gocore

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type CoreClientTcp struct {
	CertFile string
	Host     string
	Port     string
	SSL      bool
}

func NewCoreClientTcp(client *CoreClientTcp) *CoreClientTcp {
	client.handleErrorHost()
	client.handleErrorPort()
	return client
}

func (c *CoreClientTcp) Start() *grpc.ClientConn {
	var ops = grpc.WithInsecure()
	if c.SSL {
		creds, sslErr := credentials.NewClientTLSFromFile(c.CertFile, "")
		if sslErr != nil {
			log.Fatal(sslErr)
		}
		ops = grpc.WithTransportCredentials(creds)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.Host, c.Port), ops)
	//defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func (c *CoreClientTcp) handleErrorHost() {
	if c.Host == "" {
		log.Fatal("Host is required")
	}
}
func (c *CoreClientTcp) handleErrorPort() {
	if c.Port == "" {
		log.Fatal("Port is required")
	}
}
