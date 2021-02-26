package main

import (
	"context"
	"golang-grpc-protobuf/common/config"
	"golang-grpc-protobuf/common/model"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

// GaragesServer struct
type GaragesServer struct{}

// Add garage
func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserID) (*empty.Empty, error) {
	userID := param.UserID
	garage := param.Garage

	if _, ok := localStorage.List[userID]; !ok {
		localStorage.List[userID] = new(model.GarageList)
		localStorage.List[userID].List = make([]*model.Garage, 0)
	}

	localStorage.List[userID].List = append(localStorage.List[userID].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userID)

	return new(empty.Empty), nil
}

// List of garage
func (GaragesServer) List(ctx context.Context, param *model.GarageUserID) (*model.GarageList, error) {
	userID := param.UserID

	return localStorage.List[userID], nil
}

func main() {
	server := grpc.NewServer()
	var garagesServer GaragesServer
	model.RegisterGaragesServer(server, garagesServer)

	log.Println("Starting RPC server at", config.ServiceGaragePort)

	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(server.Serve(l))

}
