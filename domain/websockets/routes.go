package websockets

import (
	"github.com/labstack/echo/v4"
)

// Handler is an interface to the HTTP handler functions.
type Handler interface {
	WebsocketHandler(c echo.Context) error
}

// SetRoutes sets all of the appropriate routes to websocket handlers for the application
func SetRoutes(engine *echo.Group, h Handler) {
	ws := engine.Group("/ws")

	ws.GET("/:channel", h.WebsocketHandler)
}
