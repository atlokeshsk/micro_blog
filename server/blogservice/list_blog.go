package blogservice

import (
	"context"

	"github.com/atlokeshsk/micro_blog/models"
	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (b *BlogService) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	cursor, err := b.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error %v", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		data := &models.Blog{}
		err = cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, "Error while decoding the data %v", err)
		}
		stream.Send(data.ToProtobuf())
	}
	if err = cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error %v", err)
	}
	return nil
}
