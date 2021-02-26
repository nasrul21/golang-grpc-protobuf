package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-grpc-protobuf/common/config"
	"golang-grpc-protobuf/common/model"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return model.NewUsersClient(conn)
}

// this main function contains an example how to call grpc services
func main() {
	user1 := model.User{
		ID:       "n001",
		Name:     "Nasrul Faizin",
		Password: "b3b4s d0n9",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	user2 := model.User{
		ID:       "n002",
		Name:     "Muslimawati",
		Password: "4p4 a74laH",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	garage1 := model.Garage{
		ID:   "q001",
		Name: "Qite Qite",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.123123123,
		},
	}

	garage2 := model.Garage{
		ID:   "q001",
		Name: "Qau Qau",
		Coordinate: &model.GarageCoordinate{
			Latitude:  50.123123123,
			Longitude: 60.123123123,
		},
	}

	garage3 := model.Garage{
		ID:   "q001",
		Name: "Qamu Qamu",
		Coordinate: &model.GarageCoordinate{
			Latitude:  70.123123123,
			Longitude: 85.123123123,
		},
	}

	fmt.Println("\n", "==========> user test")
	// register user1 & user2
	userServiceTest(user1, user2)

	fmt.Println("\n", "==========> garage test A")
	// add garage1 & garage2 to user1
	garageServiceTest(user1, garage1, garage2)

	fmt.Println("\n", "==========> garage test B")
	// add garage3 to user2
	garageServiceTest(user2, garage3)
}

// call user service
func userServiceTest(users ...model.User) {
	user := serviceUser()

	// register users
	for _, u := range users {
		user.Register(context.Background(), &u)
	}

	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))
}

// call garage service
func garageServiceTest(user model.User, garages ...model.Garage) {
	garage := serviceGarage()

	// add garages to user1
	for _, g := range garages {
		garage.Add(context.Background(), &model.GarageAndUserID{
			UserID: user.ID,
			Garage: &g,
		})
	}

	// show all garages of user
	res2, err := garage.List(context.Background(), &model.GarageUserID{UserID: user.ID})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
}
