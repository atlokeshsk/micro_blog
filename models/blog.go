package models

import (
	pb "github.com/atlokeshsk/micro_blog/protogen/blog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func (b *Blog) ToProtobuf() *pb.Blog {
	return &pb.Blog{
		Id:       b.ID.Hex(),
		AuthorId: b.AuthorID,
		Title:    b.AuthorID,
		Content:  b.Content,
	}
}

func BlogFromProtobuf(in *pb.Blog) *Blog {
	return &Blog{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
}
