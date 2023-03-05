package gateway

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func handleConnection(s *Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := s.Upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		s.Sockets[uuid.New().String()] = ws

		return nil
	}
}
