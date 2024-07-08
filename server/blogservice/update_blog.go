package blogservice

import (
	"context"

	"github.com/atlokeshsk/micro_blog/models"
	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (b *BlogService) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse id %v", err)
	}
	blog := models.BlogFromProtobuf(in)
	res, err := b.Collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": blog},
	)
	if err != nil{
		return nil, status.Error(codes.Internal,"Internal Error")
	}

	if res.MatchedCount == 0{
		return nil, status.Errorf(codes.NotFound,"no document found with given id : %s",id)
	}
	return &emptypb.Empty{}, nil

}
