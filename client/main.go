package main

import (
	"context"
	"fmt"
	pb "go_grpc/helloworld"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// address     = "localhost:50051"//客户端和服务端部署在同一台服务器
	address     = "81.68.81.xx:50051" //客户端和服务端在不同服务器，需要把这里换成服务端所在服务器的IP地址
	defaultId   = 9527
	defaultName = "cowboy"
	defaultDesc = "very busy"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewUserInfoClient(conn)
	ctx := context.Background()
	id := defaultId
	if len(os.Args) > 1 {
		id, _ = strconv.Atoi(os.Args[1])
	}
	name := defaultName
	if len(os.Args) > 2 {
		name = os.Args[2]
	}
	desc := defaultDesc
	if len(os.Args) > 3 {
		desc = os.Args[3]
	}

	aMac := &pb.User{Id: uint32(id), Name: name, Description: desc}
	userDesc, err := client.GetUserDesc(ctx, aMac)
	if err != nil {
		log.Println("fail:", err)
		return
	}
	fmt.Printf("userDesc: %v\n", userDesc)
	log.Println("success:", userDesc.Desc)
}
