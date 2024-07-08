package main

import (
	"context"
	"log"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
)

func createBlog(client pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "123",
		Title:    "Amarkalam",
		Content:  "Ajith",
	}
	result, err := client.CreateBlog(context.TODO(), blog)
	if err != nil {
		log.Fatalf("failed to create blog %v", err)
	}
	log.Printf("blog created successfully id : %s", result.Id)
	return result.Id
}
