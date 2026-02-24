package main

import (
	"context"
	"go-microservices/common/config"
	"go-microservices/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (GaragesServer) List(_ context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	log.Println("Listing Garages for User", param.UserId)
	if list, ok := localStorage.List[param.UserId]; ok {
		return list, nil
	}
	return new(model.GarageList), nil
}

func (GaragesServer) Add(_ context.Context, param *model.GarageAndUserId) (*emptypb.Empty, error) {
	log.Println("Adding Garage for User", param.UserId, "Garage", param.Garage.String())	

	userId := param.UserId
    garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	return new(emptypb.Empty), nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.ServiceGaragePort)

	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(srv.Serve(l))
}