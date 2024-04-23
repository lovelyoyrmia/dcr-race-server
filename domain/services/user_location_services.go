package services

import (
	"github.com/labstack/echo/v4"
	"github.com/lovelyoyrmia/dcr_race/domain/repositories"
	"github.com/lovelyoyrmia/dcr_race/domain/websockets"
	"github.com/lovelyoyrmia/dcr_race/pkg/config"
)

// Dispatcher provides an interface to dispatch events to clients connect over websockets
type Dispatcher interface {
	Broadcast() chan *websockets.Message
}

type UserLocationServices interface {
	GetUserLocations(c echo.Context) error
	CreateUserLocationService(c echo.Context) error
}

type userLocationServices struct {
	repo       repositories.UserRepositories
	config     config.Config
	dispatcher Dispatcher
}

// Action is the payload dispatched to any clients connected over websocket
type Action struct {
	Type     string      `json:"type"`
	Category string      `json:"category"`
	Path     string      `json:"path"`
	Data     interface{} `json:"data"`
}

func NewUserLocationServices(repo repositories.UserRepositories, config config.Config, dispatcher Dispatcher) UserLocationServices {
	return &userLocationServices{
		repo:       repo,
		config:     config,
		dispatcher: dispatcher,
	}
}

func (h *userLocationServices) broadcast(typ, path, category string, b interface{}) {
	// dispatch action

	action := Action{
		Type:     typ,
		Path:     path,
		Category: category,
		Data:     b,
	}
	h.dispatcher.Broadcast() <- &websockets.Message{Channel: category, Data: action}
}
