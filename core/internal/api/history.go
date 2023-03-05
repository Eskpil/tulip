package api

import (
	"encoding/json"
	"github.com/eskpil/tulip/core/internal/api/responses"
	"github.com/eskpil/tulip/core/internal/database"
	"github.com/eskpil/tulip/core/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LastState(c echo.Context) error {
	id := c.Param("id")

	state := new(models.EntityState)
	result := database.Client().Order("created_at DESC").Find(state, "entity_id = ?", id)
	if result.Error != nil {
		log.Infof("Could not get the last state of entity: %s: %v", id, result.Error)
		return result.Error
	}

	response := new(responses.State)

	response.EntityId = state.EntityId
	response.Attributes = make(map[string]interface{})

	if err := json.Unmarshal([]byte(state.State), &response.State); err != nil {
		log.Errorf("Could not unmarshal state: %v", err)
		return err
	}

	response.UpdatedAt = state.UpdatedAt
	response.CreatedAt = state.CreatedAt

	return c.JSON(http.StatusOK, response)
}
