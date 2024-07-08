package main

import (
	"context"
	"log"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	_, err := client.DeleteBlog(context.TODO(), &pb.BlogID{Id: id})
	if err != nil {
		log.Fatalf("failed to delete the blog %v", err)
	}
	log.Println("blog Dleted")
}
