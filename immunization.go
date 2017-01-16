package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetImmunization will return a careplan for a patient with id pid
func (c *Connection) GetImmunization(pid string) (*Immunization, error) {
	b, err := c.Query(fmt.Sprintf("Immunization?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := Immunization{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Immunization is a FHIR immunization
type Immunization struct {
	ResourceType string    `json:"resourceType"`
	Status       string    `json:"status"`
	Date         time.Time `json:"date"`
	WasNotGiven  bool      `json:"wasNotGiven"`
	Reported     bool      `json:"reported"`
	LotNumber    string    `json:"lotNumber"`
	ID           string    `json:"id"`
	VaccineCode  Concept   `json:"vaccineCode"`
	Patient      Person    `json:"patient"`
}
