package repositories

import (
	"context"

	"github.com/lovelyoyrmia/dcr_race/internal/db"
	"github.com/lovelyoyrmia/dcr_race/pkg/types"
)

type UserRepositories interface {
	SaveUserLocation(ctx context.Context, req *types.UserLocationParams) (*db.UsersLocation, error)
	GetUserLocations(ctx context.Context, req *types.ListUserLocationParams) (*[]db.UsersLocation, error)
}

type userRepositories struct {
	db db.Querier
}

func NewUserRepositories(db db.Querier) UserRepositories {
	return &userRepositories{
		db: db,
	}
}
