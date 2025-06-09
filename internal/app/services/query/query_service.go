package query

import (
	"context"

	"github.com/naveenm4d/bff-gql/graph/model"
	"github.com/naveenm4d/bff-gql/internal/app/services/utils/composer"
	"github.com/naveenm4d/bff-gql/internal/core/adapters"

	blogpb "github.com/naveenm4d/blog-svc/proto"
)

var _ = adapters.QueryService(&queryService{})

type queryService struct {
	blogSvcClient blogpb.BlogSvcClient
}

func NewQueryService(
	blogSvcClient blogpb.BlogSvcClient,
) adapters.QueryService {
	service := &queryService{
		blogSvcClient: blogSvcClient,
	}

	return service
}

func (q *queryService) GetPosts(ctx context.Context) ([]*model.Post, error) {
	response, err := q.blogSvcClient.GetPosts(ctx, &blogpb.GetPostsRequest{})
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)

	for _, post := range response.GetPosts() {
		posts = append(posts, composer.GetPost(post))
	}

	return posts, nil
}
