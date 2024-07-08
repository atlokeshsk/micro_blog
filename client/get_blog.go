package main

import (
	"context"
	"log"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
)

func getBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readblog was invoked")
	blogID := &pb.BlogID{Id: id}
	result, err := client.GetBlog(context.TODO(), blogID)
	if err != nil {
		log.Fatalf("failed to get blog %v", err)
	}
	return result
}
