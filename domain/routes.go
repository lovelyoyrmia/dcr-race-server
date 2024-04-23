package domain

import (
	"github.com/labstack/echo/v4"
	"github.com/lovelyoyrmia/dcr_race/domain/services"
)

type UserLocationRoutes struct {
	services services.UserLocationServices
}

func NewUserLocationRoutes(services services.UserLocationServices) *UserLocationRoutes {
	return &UserLocationRoutes{
		services: services,
	}
}

func (route *UserLocationRoutes) SetupRoutes(c *echo.Group) error {
	c.POST("/submit", route.services.CreateUserLocationService)
	c.GET("/locations", route.services.GetUserLocations)
	return nil
}
