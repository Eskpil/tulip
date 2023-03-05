package gateway

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Server struct {
	E *echo.Echo

	Sockets []*websocket.Conn
}
