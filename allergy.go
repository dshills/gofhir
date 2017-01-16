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
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Criticality string     `json:"criticality"`
			Onset       time.Time  `json:"onset"`
			Substance   Concept    `json:"substance"`
			Reaction    []Reaction `json:"reaction"`
			Note        Note       `json:"note"`
		} `json:"resource"`
	} `json:"entry"`
}
