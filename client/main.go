package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-microservices/common/config"
	"go-microservices/common/model"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceGarage() model.GaragesClient {
	port := config.ServiceGaragePort
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.ServiceUserPort
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))	
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}

func main()	{
	user1 := model.User{
		Id: "1",
		Name: "John Doe",
		Password: "Apa lu ?",
		Gender: model.UserGender_MALE,
	}

	// garage1 := model.Garage{
	// 	Id: "1",
	// 	Name: "John's Garage",
	// 	Coordinate: &model.GarageCoordinate{
	// 		Latitude: 1.234567,
	// 		Longitude: 2.345678,
	// 	},
	// }

	user := serviceUser()

	fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))

	user.Register(context.Background(), &user1)

	// show all registered users
	res1, err := user.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

}
