package api_graphql

import (
	"context"

	graphql1 "github.com/iot-for-tillgenglighet/api-graphql/pkg/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Temperatures(ctx context.Context) ([]*graphql1.Temperature, error) {
	panic("not implemented")
}
