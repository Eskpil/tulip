package responses

import (
	"encoding/json"
	"github.com/eskpil/tulip/core/pkg/models"
	"time"
)

type State struct {
	EntityId   string                 `json:"entity_id"`
	State      interface{}            `json:"state"`
	Attributes map[string]interface{} `json:"attributes"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

func FromEntityState(entityState models.EntityState) (*State, error) {
	state := new(State)

	state.EntityId = entityState.EntityId
	state.Attributes = make(map[string]interface{})

	if err := json.Unmarshal([]byte(entityState.State), &state.State); err != nil {
		return nil, err
	}

	state.UpdatedAt = entityState.UpdatedAt
	state.CreatedAt = entityState.CreatedAt

	return state, nil
}

func (s *State) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(s);
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
