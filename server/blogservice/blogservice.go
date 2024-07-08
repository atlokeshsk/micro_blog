package blogservice

import (
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
)

type BlogService struct {
	pb.UnimplementedBlogServiceServer
	Collection *mongo.Collection
}
