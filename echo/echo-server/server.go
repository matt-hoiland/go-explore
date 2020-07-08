package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/matt-hoiland/go-explore/echo/echopb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type customEchoServiceServer struct{}

func (*customEchoServiceServer) EchoOnce(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	log.Debugln("customEchoServiceServer.EchoOnce invoked")
	log.Debugf("Context received: %v\n", ctx)
	log.Debugf("Request received: %v\n", req)

	return &echopb.EchoResponse{
		Echo: req.GetMessage(),
	}, nil
}

func (*customEchoServiceServer) EchoMultiple(req *echopb.EchoRequest, stream echopb.EchoService_EchoMultipleServer) error {
	log.Debugln("customEchoServiceServer.EchoMultiple invoked")
	log.Debugf("Request received: %v\n", req)

	message := []rune(req.GetMessage())
	log.Debugf("message length: %d\n", len(message))

	for i := len(message); i >= 0; i-- {
		response := string(message[0:i])
		if i < len(message) {
			response += "..."
		}

		if err := stream.Send(&echopb.EchoResponse{Echo: response}); err != nil {
			log.Fatalf("Failed to send response: %v\n", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

var (
	// PORT is the port number for the grpc server
	PORT = 50051
	// HOST is the host of the grpc server
	HOST = "0.0.0.0"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	ADDR := fmt.Sprintf("%s:%d", HOST, PORT)

	lis, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Infof("Address %s reserved\n", ADDR)

	s := grpc.NewServer()
	echopb.RegisterEchoServiceServer(s, &customEchoServiceServer{})

	log.Infoln("Serving [ctrl-c to stop]...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
