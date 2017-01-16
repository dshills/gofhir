package fhir

import (
	"encoding/json"
	"fmt"
)

// TODO more types of queries

// GetMedication will return a careplan for a patient with id pid
func (c *Connection) GetMedication(pid string) (*Medication, error) {
	b, err := c.Query(fmt.Sprintf("MedicationOrder?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := Medication{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Medication is a FHIR medication
type Medication struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			DateWritten         string              `json:"dateWritten"`
			Identifier          []Identifier        `json:"identifier"`
			Prescriber          Person              `json:"prescriber"`
			MedicationReference Thing               `json:"medicationReference"`
			DosageInstruction   []DosageInstruction `json:"dosageInstruction"`
			DispenseRequest     DispenseRequest     `json:"dispenseRequest"`
		} `json:"resource"`
	} `json:"entry"`
}
