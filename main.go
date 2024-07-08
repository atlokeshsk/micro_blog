package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"github.com/atlokeshsk/micro_blog/server/blogservice"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file %v", err)
	}
	client := openDB()
	defer closeDB(client)
	log.Println("server connected to the mongodb successfully")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatalf("failed to get PORT from .env file")
	}
	address := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listend on %s %v", port, err)
	}
	grpcServer := grpc.NewServer()
	blogService := &blogservice.BlogService{Collection: client.Database("micro_blog").Collection("blog")}
	pb.RegisterBlogServiceServer(grpcServer, blogService)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start server %v", err)
	}
}

func openDB() *mongo.Client {
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatalf("MONGODB_URI env variable not set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
	return client
}
func closeDB(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
