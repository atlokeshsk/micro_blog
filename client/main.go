package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("failed to load .env file %v", err)
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("failed to get PORT from .env file")
	}
	address := fmt.Sprintf(":%s", port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to make connection %v", err)
	}
	blogClient := pb.NewBlogServiceClient(conn)
	id := createBlog(blogClient)
	log.Println(getBlog(blogClient, id))
	updateBlog(blogClient, id)
	listBlog(blogClient)
	deleteBlog(blogClient, id)
}
