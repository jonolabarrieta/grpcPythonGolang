package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	pb "google.golang.org/grpc/examples/routeguide"
	"google.golang.org/grpc/reflection"
)

var (
	log  hclog.Logger
	port = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {

	go func() {
		for {
			rr, err := stream.Recv()
			if err == io.EOF {
				log.Error("Client write closed")
				break
			}

			if err != nil && err != io.EOF {
				log.Error("Error reading from client", "error", err)
				break
			}

			log.Info("New message from client", "Location", rr.GetLocation(), "Message", rr.GetMessage())
		}
	}()

	for {
		log.Info("Send message to client")
		time.Sleep(1 * time.Second)

		data := &pb.RouteNote{Location: "server", Message: "Desde el server"}
		err := stream.Send(data)
		if err != nil {
			log.Error("Error sending to client", "error", err)
			break
		}
	}

	return nil
}

func main() {
	log = hclog.Default()
	fmt.Println("Server listenig on -->localhost:", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Error("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &routeGuideServer{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}

}
