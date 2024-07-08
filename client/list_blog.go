package main

import (
	"context"
	"io"
	"log"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(client pb.BlogServiceClient) {
	log.Println("list blog invoked")
	stream, err := client.ListBlog(context.TODO(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error while calling list blog %v", err)
	}
	for {
		pbBlog, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to list blog %v", err)
		}
		log.Println(pbBlog)
	}
}
