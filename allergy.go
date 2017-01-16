package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetAllergyIntolerence will return patient allergy intolerence
func (c *Connection) GetAllergyIntolerence(pid string) (*AllergyIntolerence, error) {
	b, err := c.Query(fmt.Sprintf("AllergyIntolerance?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := AllergyIntolerence{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// AllergyIntolerence is a FHIR allergy intolerence
type AllergyIntolerence struct {
	ResourceType string `json:"resourceType"`
	Type         string `json:"type"`
	Total        int    `json:"total"`
	Link         []Link `json:"link"`
	Entry        []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Criticality string    `json:"criticality"`
			Onset       time.Time `json:"onset"`
			Recorder    Person    `json:"recorder"`
			Substance   Concept   `json:"substance"`
			Note        Note      `json:"note"`
			Reaction    Reaction  `json:"reaction"`
		} `json:"resource"`
	} `json:"entry"`
}
