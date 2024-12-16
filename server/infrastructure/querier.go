// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package infrastructure

import (
	"context"
)

type Querier interface {
	FindAllRestaurants(ctx context.Context) ([]Restaurant, error)
	FindAllTemporaryUsers(ctx context.Context) ([]TemporaryUser, error)
	FindAllUsers(ctx context.Context) ([]User, error)
	FindAllVisitedRestaurants(ctx context.Context) ([]VisitedRestaurant, error)
}

var _ Querier = (*Queries)(nil)