package main

import (
	"context"
	"fmt"
	pb "go_grpc/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserInfoServer
}

const (
	port = ":50051"
)

func (s *server) GetUserDesc(cxt context.Context, user *pb.User) (*pb.UserDesc, error) {
	desc := fmt.Sprintf("用户编号是:%d,名称是:%s,描述是:%s", user.Id, user.Name, user.Description)
	resp := &pb.UserDesc{
		Desc: desc,
	}
	return resp, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &server{})
	log.Println("start gRPC listen on port " + port)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
