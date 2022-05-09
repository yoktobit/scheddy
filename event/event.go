package event

import (
	"golang.org/x/net/context"
)

type eventServiceServer struct {
	UnimplementedEventServiceServer
}

func (es *eventServiceServer) NewEvent(context context.Context, newEventRequest *NewEventRequest) (*NewEventResponse, error) {
	return &NewEventResponse{
		Id:   "1",
		Name: newEventRequest.Name,
	}, nil
}

func NewEventServiceServer() EventServiceServer {

	return &eventServiceServer{}
}
