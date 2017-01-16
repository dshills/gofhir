package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetDiagnosticReport will return a diagnostic report for a patient with id pid
func (c *Connection) GetDiagnosticReport(pid string) (*DiagnosticReport, error) {
	b, err := c.Query(fmt.Sprintf("DiagnosticReport?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := DiagnosticReport{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DiagnosticReport is a FHIR report
type DiagnosticReport struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			EffectiveDateTime time.Time    `json:"effectiveDateTime"`
			Issued            time.Time    `json:"issued"`
			Identifier        []Identifier `json:"identifier"`
			Code              Note         `json:"code"`
			Result            []Thing      `json:"result"`
		} `json:"resource"`
	} `json:"entry"`
}
