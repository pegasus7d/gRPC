package main

import (
	"context"
	"log"

	pb "github.com/pegasus7d/grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient){
	log.Println("DoSum was invoked")
	res,err:=c.Sum(context.Background(),&pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 1,
	})

	if err!=nil{
		log.Fatalf("Could not calculate:%v\n",err)
	}

	log.Printf("Sum: %d\n",res.Result)

}