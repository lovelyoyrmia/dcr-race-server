package repositories

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lovelyoyrmia/dcr_race/internal/db"
	"github.com/lovelyoyrmia/dcr_race/pkg/types"
)

// SaveUserLocation implements UserRepositories.
func (u *userRepositories) SaveUserLocation(ctx context.Context, req *types.UserLocationParams) (*db.UsersLocation, error) {

	user, err := u.db.GetUserLocation(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			result, err := u.db.CreateUserLocation(ctx, db.CreateUserLocationParams{
				Uid: req.Uid,
				Email: sql.NullString{
					String: req.Email,
					Valid: true,
				},
				Latitude: sql.NullFloat64{
					Float64: req.Latitude,
					Valid:   true,
				},
				Longitude: sql.NullFloat64{
					Float64: req.Longitude,
					Valid:   true,
				},
				Altitude: sql.NullFloat64{
					Float64: req.Altitude,
					Valid:   true,
				},
				Category: req.Category,
				Fullname: req.Fullname,
			})

			if err != nil {
				return nil, err
			}

			uid, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}

			log.Println(uid)

			newUser, err := u.db.GetUserLocationByID(ctx, int32(uid))
			if err != nil {
				return nil, err
			}
			return &newUser, nil
		}

		return nil, err
	}

	_, err = u.db.UpdateUserLocation(ctx, db.UpdateUserLocationParams{
		Uid: user.Uid,
		Email: sql.NullString{
			String: req.Email,
			Valid: true,
		},
		Fullname: req.Fullname,
		Latitude: sql.NullFloat64{
			Float64: req.Latitude,
			Valid:   true,
		},
		Longitude: sql.NullFloat64{
			Float64: req.Longitude,
			Valid:   true,
		},
		Altitude: sql.NullFloat64{
			Float64: req.Altitude,
			Valid:   true,
		},
		Timestamp: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	newUser, err := u.db.GetUserLocationByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &newUser, nil

}
