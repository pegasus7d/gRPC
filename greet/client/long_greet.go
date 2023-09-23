package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pegasus7d/grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet was invoked")
	reqs:=[]*pb.GreetRequest{
		{FirstName: "Debayan"},
		{FirstName: "Eshanika"},
		{FirstName: "Debayan"},
		{FirstName: "Eshanika"},
	}
	stream,err:=c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _,req:= range reqs{
		log.Printf("Sending req:%v\n",req)
		stream.Send(req)
		time.Sleep(5*time.Second)
	}

	res,err:=stream.CloseAndRecv()

	if err!=nil{
		log.Fatalf("Error while receiving response%v\n",err)
	}

	log.Printf("LongGreet %s",res.Result)

}