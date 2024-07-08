package blogservice

import (
	"context"

	"github.com/atlokeshsk/micro_blog/models"
	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogID, error) {
	blog := models.BlogFromProtobuf(in)
	result, err := b.Collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error %v", err)
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {

		return nil, status.Error(codes.Internal, "failed to convert the id to object id")
	}
	return &pb.BlogID{
		Id: id.Hex(),
	}, nil
}
