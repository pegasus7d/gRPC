package main

import (
	"context"
	"log"

	pb "github.com/pegasus7d/grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("DoGreet was invoked")
	res,err:=c.Greet(context.Background(),&pb.GreetRequest{
		FirstName: "Debayan",
	})

	if err!=nil{
		log.Fatalf("Could not greet:%v\n",err)
	}

	log.Printf("Greeting: %s\n",res.Result)

}