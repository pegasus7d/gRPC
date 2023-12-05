package main

import (
	"context"
	"log"

	pb "github.com/pegasus7d/grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient)string{
	log.Println("---createBlog was invoked---")
	blog:=&pb.Blog{
		AuthorId:"Eshanika",
		Title: "One of My Fav",
		Content: "Content of One of My Fav",
	}
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
		return ""
	}

	log.Printf("Blog has been created: %v\n", res.Id)
	return res.Id
}