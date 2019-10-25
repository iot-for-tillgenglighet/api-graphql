package api_graphql

import (
	"context"

	graphql1 "github.com/iot-for-tillgenglighet/api-graphql/pkg/graphql"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	temperatures []*graphql1.Temperature
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Temperatures(ctx context.Context) ([]*graphql1.Temperature, error) {
	if len(r.temperatures) == 0 {
		temp := &graphql1.Temperature{
			From: graphql1.Origin{
				Device: "mydevice",
				Lat:    12.3,
				Lon:    16.8,
			},
			When: "1900-01-01T12:00:00",
			Temp: 37.0,
		}
		r.temperatures = append(r.temperatures, temp)
	}
	return r.temperatures, nil
}
