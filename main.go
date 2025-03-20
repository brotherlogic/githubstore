package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	github_client "github.com/brotherlogic/githubridge/client"

	pb "github.com/brotherlogic/githubridge/proto"
)

var (
	port        = flag.Int("port", 8080, "Server port for grpc traffic")
	metricsPort = flag.Int("metrics_port", 8081, "Metrics port")
)

type Server struct {
	client github_client.GithubridgeClient
}

func NewServer() *Server {
	client, err := github_client.GetClientInternal()

	if err != nil {
		panic(err)
	}

	return &Server{
		client: client,
	}
}

func main() {
	flag.Parse()

	s := NewServer()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%v", *metricsPort), nil)
		log.Fatalf("gramophile is unable to serve metrics: %v", err)
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("mdb is unable to listen on the internal grpc port %v: %v", *port, err)
	}
	gsInt := grpc.NewServer()
	pb.RegisterGithubridgeServiceServer(gsInt, s)

	err = gsInt.Serve(lis)
	log.Fatalf("mdb is unable to sever grpc internally: %v", err)

}
