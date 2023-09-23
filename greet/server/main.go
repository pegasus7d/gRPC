package main

import (
	"log"
	"net"

	pb "github.com/pegasus7d/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

type Server struct{
	pb.GreetServiceServer
}

func main() {

	lis,err := net.Listen("tcp",addr)
	if err!=nil{
		log.Fatalf("Failed to listen on:%v\n",err)
	}

	log.Printf("Listening on %s\n",addr)
	opts:=[]grpc.ServerOption{}
	tls:=true 
	if tls{
		certFile:="ssl/server.crt"
		keyFile:="ssl/server.pem"
		creds,err:=credentials.NewServerTLSFromFile(certFile,keyFile)

		if err!=nil{
			log.Fatalf("Failed Loading cerificates%v\n",err)
		}
		opts=append(opts, grpc.Creds(creds))
		
	}
	s:=grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s,&Server{
		
	})

	if err=s.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve on:%v\n",err)
	}
}