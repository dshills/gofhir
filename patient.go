package fhir

import (
	"encoding/json"
	"fmt"
)

// GetPatient will return patient information for a patient with id pid
func (c *Connection) GetPatient(pid string) (*Patient, error) {
	b, err := c.Query(fmt.Sprintf("Patient/%v", pid))
	if err != nil {
		return nil, err
	}
	data := Patient{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Patient is a FHIR patient
type Patient struct {
	ResourceType    string          `json:"resourceType"`
	BirthDate       string          `json:"birthDate"`
	Active          bool            `json:"active"`
	Gender          string          `json:"gender"`
	DeceasedBoolean bool            `json:"deceasedBoolean"`
	ID              string          `json:"id"`
	CareProvider    []Person        `json:"careProvider"`
	Name            []Name          `json:"name"`
	Identifier      []Identifier    `json:"identifier"`
	Address         []Address       `json:"address"`
	Telecom         []Telecom       `json:"telecom"`
	MaritalStatus   Concept         `json:"maritalStatus"`
	Communication   []Communication `json:"communication"`
	Extension       []Extension     `json:"extension"`
}
