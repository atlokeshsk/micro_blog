package main

import (
	"context"
	"log"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
)

func updateBlog(client pb.BlogServiceClient, id string) {
	log.Println("update blog invoked")

	pbBlog := &pb.Blog{
		Id:       id,
		AuthorId: "updated",
		Title:    "ajith",
		Content:  "update",
	}

	_, err := client.UpdateBlog(context.TODO(), pbBlog)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

}
