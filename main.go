package main

import (
	"context"
	pb "github.com/oyamoh-brian/tv-service-database/proto/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	rpcPort = ":50050"
)


type service struct {
	pb.DataBaseServiceServer
}

func (s *service) GetAllVideos(context.Context, *pb.Category) (*pb.ResponseVideos, error)  {
	var resp = pb.ResponseVideos{}
	resp.Status = 200
	resp.Message = "Success"

	resp.Video = *GetDummyVideos(nil)

	return &resp, nil
}

func (s *service)GetAllCategories(context.Context, *pb.GetRequest) (*pb.ResponseCategories, error)  {
	var resp = pb.ResponseCategories{}
	resp.Status = 200
	resp.Message = "Success"

	resp.Categories = *GetDummyCategories()

	return  &resp, nil
}

func main()  {


	lis, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterDataBaseServiceServer(s, &service{})

	reflection.Register(s)

	log.Println("Database service running on port:", rpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}