package composer

import (
	"github.com/google/uuid"
	"github.com/naveenm4d/bff-gql/graph/model"
	blogpb "github.com/naveenm4d/blog-svc/proto"
)

func GetPost(post *blogpb.Post) *model.Post {
	return &model.Post{
		ID:       uuid.New(),
		AuthorID: uuid.New(),
		Slug:     post.GetSlug(),
		Title:    post.GetTitle(),
		Content:  post.GetContent(),
		Status:   model.PostStatus(post.Status.String()),
	}
}
