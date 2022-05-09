package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/yoktobit/scheddy_backend/event"
	"google.golang.org/grpc"
)

func main() {
	//dataaccess.NewConnectionWithEnvironment()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listener could not be run")
	}
	eventServiceServer := event.NewEventServiceServer()
	grpcServer := grpc.NewServer()
	event.RegisterEventServiceServer(grpcServer, eventServiceServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Server could not be started")
	}
}
