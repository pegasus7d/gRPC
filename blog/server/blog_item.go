package main

import (
	pb "github.com/pegasus7d/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type BlogItem struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Author string `bson:"author_id"`
	Title string `bson:"title"`
	Content string `bson:"context"`
}
func documentToBlog(data *BlogItem)*pb.Blog{
	return &pb.Blog{
		Id: data.ID.Hex(),
		AuthorId: data.Author,
		Title: data.Title,
		Content: data.Content,
	}
}