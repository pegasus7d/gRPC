package main

import (
	"log"
	"net"

	pb "github.com/pegasus7d/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "localhost:50051"

type Server struct{
	pb.CalculatorServiceServer
}

func main() {

	lis,err := net.Listen("tcp",addr)
	if err!=nil{
		log.Fatalf("Failed to listen on:%v\n",err)
	}

	log.Printf("Listening on %s\n",addr)
	s:=grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s,&Server{
	
	})
	reflection.Register(s)

	if err=s.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve on:%v\n",err)
	}
}