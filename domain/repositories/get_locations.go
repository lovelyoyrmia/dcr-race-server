package repositories

import (
	"context"

	"github.com/lovelyoyrmia/dcr_race/internal/db"
	"github.com/lovelyoyrmia/dcr_race/pkg/types"
)

// GetUserLocations implements UserRepositories.
func (u *userRepositories) GetUserLocations(ctx context.Context, req *types.ListUserLocationParams) (*[]db.UsersLocation, error) {
	var limit int64 = 1000
	if req.Count != nil {
		limit = *req.Count
	}

	if req.Category == nil {
		users, err := u.db.GetUserLocations(ctx, int32(limit))
		if err != nil {
			return nil, err
		}

		return &users, nil
	}

	users, err := u.db.GetUserLocationsByCategory(ctx, db.GetUserLocationsByCategoryParams{
		Category: *req.Category,
		Limit:    int32(limit),
	})
	if err != nil {
		return nil, err
	}

	return &users, nil
}
