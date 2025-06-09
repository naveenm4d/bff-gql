package graph

import "github.com/naveenm4d/bff-gql/internal/core/adapters"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	QueryService adapters.QueryService
}
