package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lovelyoyrmia/dcr_race/pkg/types"
)

func (service *userLocationServices) GetUserLocations(c echo.Context) error {
	m := new(types.ListUserLocationParams)

	if err := c.Bind(m); err != nil {
		errorRes := map[string]string{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorRes)
	}

	locations, err := service.repo.GetUserLocations(c.Request().Context(), m)

	if err != nil {
		errorRes := map[string]string{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorRes)
	}

	userLocations := make([]types.UserLocation, 0)

	for _, v := range *locations {
		userLocations = append(userLocations, types.UserLocation{
			Uid:       v.Uid,
			Email:     v.Email.String,
			Fullname:  v.Fullname,
			Latitude:  v.Latitude.Float64,
			Longitude: v.Longitude.Float64,
			Altitude:  v.Altitude.Float64,
			Category:  v.Category,
			Timestamp: v.Timestamp.Time.String(),
		})
	}

	return c.JSON(http.StatusOK, userLocations)
}
