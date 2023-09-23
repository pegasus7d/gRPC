package main

import (
	"io"
	"log"

	pb "github.com/pegasus7d/grpc/calculator/proto"
)

func (s *Server)Max(stream pb.CalculatorService_MaxServer) error{
	log.Println("Max function was invoked")
	var maximum int64=0


	for{
		req,err:=stream.Recv()
		if err==io.EOF{
			return nil
		}
		if err!=nil{
			log.Printf("Error while reading client stream %v\n",err)
		}

		if number:=req.Number;number>int64(maximum){
			maximum=number
			err:=stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err!=nil{
				log.Fatalf("Error while sending data to client :%v\n",err)
			}
			
		}
	}



}