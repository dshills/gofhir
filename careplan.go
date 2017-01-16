package fhir

import (
	"encoding/json"
	"fmt"
)

// GetCarePlan will return a careplan for a patient with id pid
func (c *Connection) GetCarePlan(pid string) (*CarePlan, error) {
	b, err := c.Query(fmt.Sprintf("CarePlan?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := CarePlan{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// CarePlan is a FHIR care plan
type CarePlan struct {
	ResourceType string `json:"resourceType"`
	Type         string `json:"type"`
	Total        int    `json:"total"`
	Link         []Link `json:"link"`
	Entry        []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Addresses []Address `json:"addresses"`
			Goal      []Thing   `json:"goal"`
			Activity  []struct {
				Detail struct {
					Prohibited bool    `json:"prohibited"`
					Category   Concept `json:"category"`
					Code       Note    `json:"code"`
				} `json:"detail"`
			} `json:"activity"`
			Category []Concept `json:"category"`
		} `json:"resource"`
	} `json:"entry"`
}
