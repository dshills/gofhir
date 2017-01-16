package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetObservation will return a careplan for a patient with id pid
func (c *Connection) GetObservation(pid string) (*Observation, error) {
	b, err := c.Query(fmt.Sprintf("Observation?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := Observation{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Observation is a FHIR observation
type Observation struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourceType      string    `json:"resourceType"`
			EffectiveDateTime time.Time `json:"effectiveDateTime"`
			Status            string    `json:"status"`
			ID                string    `json:"id"`
			Code              Concept   `json:"code"`
			ValueQuantity     Quantity  `json:"valueQuantity"`
			Subject           Thing     `json:"subject"`
			Performer         []Person  `json:"performer"`
			Category          Concept   `json:"category"`
		} `json:"resource"`
	} `json:"entry"`
}
