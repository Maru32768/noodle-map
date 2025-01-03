// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package infra

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	FindAllCategories(ctx context.Context) ([]FindAllCategoriesRow, error)
	FindAllRestaurants(ctx context.Context) ([]FindAllRestaurantsRow, error)
	FindAllTemporaryUsers(ctx context.Context) ([]TemporaryUser, error)
	FindCategoriesByRestaurantIds(ctx context.Context, dollar_1 []uuid.UUID) ([]FindCategoriesByRestaurantIdsRow, error)
}

var _ Querier = (*Queries)(nil)
