package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetDevice will return a device for a patient with id pid
func (c *Connection) GetDevice(pid string) (*Device, error) {
	b, err := c.Query(fmt.Sprintf("Device?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := Device{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Device is a FHIR device
type Device struct {
	ResourceType string `json:"resourceType"`
	Type         string `json:"type"`
	Total        int    `json:"total"`
	Link         []Link `json:"link"`
	Entry        []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			UDI    string    `json:"udi"`
			Model  string    `json:"model"`
			Expiry time.Time `json:"expiry"`
			Type   Concept   `json:"coding"`
		} `json:"resource"`
	} `json:"entry"`
}
