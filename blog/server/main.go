package main

import (
	"context"
	"log"
	"net"

	pb "github.com/pegasus7d/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)


var addr string = "localhost:50051"

type Server struct{
	pb.BlogServiceServer
}

func main() {
	client,err:=mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@mongo:27017/?directConnection=true"))
	//mongodb://user:password@mongo:27017/?connect=direct
	if err!=nil{
		log.Fatal(err)
	}
	collection=client.Database("blogdb").Collection("blog")
	err=client.Connect(context.Background())
	lis,err := net.Listen("tcp",addr)
	if err!=nil{
		log.Fatalf("Failed to listen on:%v\n",err)
	}

	log.Printf("Listening on %s\n",addr)
	s:=grpc.NewServer()

	pb.RegisterBlogServiceServer(s,&Server{
	
	})
	

	if err=s.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve on:%v\n",err)
	}
}