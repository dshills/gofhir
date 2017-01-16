package fhir

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetDocumentReference will return a careplan for a patient with id pid
func (c *Connection) GetDocumentReference(pid string) (*DocumentReference, error) {
	b, err := c.Query(fmt.Sprintf("DocumentReference?patient=%v", pid))
	if err != nil {
		return nil, err
	}
	data := DocumentReference{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// DocumentReference is a FHIR document
type DocumentReference struct {
	SearchResult
	Entry []struct {
		EntryPartial
		Resource struct {
			ResourcePartial
			Created          time.Time    `json:"created"`
			Indexed          time.Time    `json:"indexed"`
			Class            Concept      `json:"class"`
			Type             Concept      `json:"type"`
			Content          []Attachment `json:"content"`
			MasterIdentifier Identifier   `json:"masterIdentifier"`
		} `json:"resource"`
	} `json:"entry"`
}
