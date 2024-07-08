package blogservice

import (
	"context"

	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (b *BlogService) DeleteBlog(ctx context.Context, in *pb.BlogID) (*emptypb.Empty, error) {

	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id %v", err)
	}
	result, err := b.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error %v", err)
	}

	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "No Document found")

	}
	return &emptypb.Empty{},nil
}
