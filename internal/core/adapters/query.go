package adapters

import (
	"context"

	"github.com/naveenm4d/bff-gql/graph/model"
)

type QueryService interface {
	GetPosts(ctx context.Context) ([]*model.Post, error)
}
