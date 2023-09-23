package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/pegasus7d/grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Println("doGreetEveryone was invoked")
	stream,err:=c.GreetEveryone(context.Background())

	if err!=nil{
		log.Fatalf("error while creating stream%v\n",err)

	}
	reqs:=[]*pb.GreetRequest{
		{FirstName: "Debayan"},
		{FirstName: "Eshanika"},
		{FirstName: "Debayan"},
		{FirstName: "Eshanika"},
	}

	waitc :=make(chan struct{})


	go func(){
		for _,req:=range reqs{
			log.Printf("Send request:%v\n",req)
			stream.Send(req)
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res,err:=stream.Recv()


			if err==io.EOF{
				// log.Printf("Error while receiving:%v\n",err)
				break;
			}
			log.Printf("Received %v\n",res.Result)
		}
		close(waitc)
	}()

	<-waitc
}

