package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lovelyoyrmia/dcr_race/pkg/types"
)

func (service *userLocationServices) CreateUserLocationService(c echo.Context) error {
	m := new(types.UserLocationParams)

	if err := c.Bind(m); err != nil {
		errorRes := map[string]string{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorRes)
	}

	userLocation, err := service.repo.SaveUserLocation(c.Request().Context(), m)

	if err != nil {
		errorRes := map[string]string{
			"message": err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errorRes)
	}

	location := types.UserLocation{
		Uid:       userLocation.Uid,
		Email:     userLocation.Email.String,
		Fullname:  userLocation.Fullname,
		Latitude:  userLocation.Latitude.Float64,
		Longitude: userLocation.Longitude.Float64,
		Altitude:  userLocation.Altitude.Float64,
		Category:  userLocation.Category,
		Timestamp: userLocation.Timestamp.Time.String(),
	}

	// dispatch action
	service.broadcast("POST", "/category", m.Category, userLocation)

	return c.JSON(http.StatusOK, location)
}
