package gocore

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
)

type CoreServerTcp struct {
	Schema   string
	Address  string
	Port     string
	Tag      string
	KeyFile  string
	CertFile string
	lis      net.Listener
	Server   *grpc.Server
}

func NewServerTcp(server *CoreServerTcp) *CoreServerTcp {
	if server.Tag == "" {
		server.Tag = "default"
	}

	server.handleErrorSchema()
	server.handleErrorAddress()
	server.handleErrorPort()
	return server.start()
}
func (s *CoreServerTcp) handleErrorSchema() {
	schema := []string{
		"http",
		"https",
		"ws",
	}
	if !slices.Contains(schema, s.Schema) {
		log.Fatal("Schema is invalid")
	}
}
func (s *CoreServerTcp) handleErrorAddress() {
	if s.Address == "" {
		log.Fatal("Host is required")
	}
}
func (s *CoreServerTcp) handleErrorPort() {
	if s.Port == "" {
		log.Fatal("port is required")
	}
}
func (s *CoreServerTcp) start() *CoreServerTcp {
	log.Printf("[TCP] Server %s running... %s://%s:%s", s.Tag, s.Schema, s.Address, s.Port)
	s.lis, _ = net.Listen("tcp", fmt.Sprintf("%s:%d", s.Address, s.Port))
	if s.CertFile == "" {
		s.CertFile = "../ssl/server.crt"
	}
	if s.KeyFile == "" {
		s.KeyFile = "../ssl/server.key"
	}
	creds, sslErr := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
	if sslErr != nil {
		log.Fatalf("create creds ssl err %v\n", sslErr)
	}

	s.Server = grpc.NewServer(grpc.Creds(creds))
	return s
}
func (s *CoreServerTcp) Run() {
	err := s.Server.Serve(s.lis)
	log.Println(err)
	//if err := s.Server.Serve(s.lis); err != nil {
	//	log.Fatal(err)
	//}
}

type ServerGateWay struct {
	Schema    string
	Address   string
	Port      string
	CertFile  string
	MuxServer *runtime.ServeMux
	Opts      []grpc.DialOption
	Ctx       context.Context
	MuxRouter *mux.Router
}

func NewServerGateWay(server *ServerGateWay) *ServerGateWay {
	server.handleErrorSchema()
	server.handleErrorAddress()
	server.handleErrorPort()
	return server.start()
}
func (s *ServerGateWay) handleErrorSchema() {
	schema := []string{
		"http",
		"https",
		"ws",
	}
	if !slices.Contains(schema, s.Schema) {
		log.Fatal("Schema is invalid")
	}
}
func (s *ServerGateWay) handleErrorAddress() {
	if s.Address == "" {
		log.Fatal("Host is required")
	}
}
func (s *ServerGateWay) handleErrorPort() {
	if s.Port == "" {
		log.Fatal("port is required")
	}
}
func (s *ServerGateWay) start() *ServerGateWay {
	s.MuxServer = runtime.NewServeMux()
	if s.CertFile == "" {
		s.CertFile = "../grpc/ssl/server.crt"
	}
	creds, sslErr := credentials.NewClientTLSFromFile(s.CertFile, s.Address)
	if sslErr != nil {
		log.Fatalf("create client creds ssl err %v\n", sslErr)
	}
	s.Opts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	r := mux.NewRouter()
	r.Handle("/", s.MuxServer)
	s.MuxRouter = r
	return s
}
func (s *ServerGateWay) Run() {
	flag.Parse()
	defer glog.Flush()
	log.Printf("[HTTP] Server gateway running... %s://%s:%s", s.Schema, s.Address, s.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.Port), s.MuxRouter); err != nil {
		glog.Fatal(err)
	}
}
