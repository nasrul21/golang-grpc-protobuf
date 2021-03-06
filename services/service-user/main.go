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

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

// UsersServer struct
type UsersServer struct{}

// Register user
func (UsersServer) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}

// List of user
func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	server := grpc.NewServer()
	var userServer UsersServer
	model.RegisterUsersServer(server, userServer)

	log.Println("Starting RPC server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}

	log.Fatal(server.Serve(l))
}
