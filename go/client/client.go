package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/routeguide"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func runRouteChat(client pb.RouteGuideClient) {
	pid := os.Getpid()
	stream, err := client.RouteChat(context.Background())
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Println("Got message at point", in.GetMessage(), in.GetMessage())
		}
	}()
	for {

		if err := stream.Send(&pb.RouteNote{Location: "locationCliente", Message: fmt.Sprint(pid)}); err != nil {
			log.Println("Failed to send a note:", err)
			stream.CloseSend()
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	runRouteChat(client)
}
