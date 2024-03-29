package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eskpil/tulip/core/internal/api/responses"
	"github.com/eskpil/tulip/core/internal/controller"
	"github.com/eskpil/tulip/core/internal/database"
	"github.com/eskpil/tulip/core/pkg/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetAll(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	_ = ctx

	var entities []models.Entity

	result := database.Client().Find(&entities)
	if result.Error != nil {
		log.Errorf("Could not get all enitites: %v", result.Error)
		return result.Error
	}

	var response []responses.Entity

	for _, e := range entities {
		r := responses.Entity{}

		r.ID = e.ID
		r.Driver = e.Driver

		r.DeviceId = e.DeviceId

		if err := json.Unmarshal([]byte(e.EntityMetadata), &r.EntityMetadata); err != nil {
			log.Errorf("Could not unmarshal entityMetadata: %v", err)
			return err
		}

		if err := json.Unmarshal([]byte(e.DriverMetadata), &r.DriverMetadata); err != nil {
			log.Errorf("Could not unmarshal driverMetadata: %v", err)
			return err
		}

		r.Name = e.Name
		r.Kind = e.Kind

		var history []models.EntityState

		result = database.Client().Limit(10).Order("created_at DESC").Find(&history, "entity_id = ?", e.ID)
		if result.Error != nil {
			log.Infof("Could not get entity history: %v", result.Error)
			return result.Error
		}

		for _, h := range history {
			state := new(responses.State)

			state.EntityId = h.EntityId
			state.Attributes = make(map[string]interface{})
			state.UpdatedAt = h.UpdatedAt
			state.CreatedAt = h.CreatedAt

			if err := json.Unmarshal([]byte(h.State), &state.State); err != nil {
				log.Errorf("Could not unmarshal state: %v", err)
				return err
			}

			r.History = append(r.History, state)
		}

		response = append(response, r)
	}

	return c.JSON(http.StatusOK, response)
}

func EntityAction(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var command controller.Command

	if err := c.Bind(&command); err != nil {
		log.Errorf("Could not bind body of request: %v", err)
		return err
	}

	id := c.Param("id")
	log.Infof("id = %s", id)

	var entity models.Entity

	result := database.Client().Find(&entity, "id = ?", id)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatalf("Could not find entity: %s", id)
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := controller.Action(ctx, &entity, &command); err != nil {
		log.Errorf("Could not perform command: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, responses.ActionResponse{Ok: true})
}
