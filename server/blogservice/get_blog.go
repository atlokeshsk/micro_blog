package blogservice

import (
	"context"

	"github.com/atlokeshsk/micro_blog/models"
	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) GetBlog(ctx context.Context, in *pb.BlogID) (*pb.Blog, error) {
	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id %v", err)
	}
	blog := &models.Blog{}
	filter := bson.M{"_id": id}
	result := b.Collection.FindOne(ctx, filter)
	if err = result.Decode(blog); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "cannot find blog with given %s", in.Id)
		}
		return nil, status.Errorf(codes.Internal, "Internal server error %v", err)
	}

	return blog.ToProtobuf(), nil
}
