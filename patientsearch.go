package fhir

import (
	"encoding/json"
	"fmt"
)

// PatientSearch will search for a patient based on the query string
// identifier, family, given, birthdate, gender, address, telecom
// i.e. family=Argonaut&given=Jason
func (c *Connection) PatientSearch(query string) (*PatientResult, error) {
	b, err := c.Query(fmt.Sprintf("Patient?%v", query))
	if err != nil {
		return nil, err
	}
	data := PatientResult{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// PatientResult is a patient search result
type PatientResult struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Patient Patient `json:"resource"`
	} `json:"entry"`
}
