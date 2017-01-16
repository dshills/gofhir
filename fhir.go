package fhir

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RetData is the mapped json of the request
type RetData map[string]interface{}

// Connection is a FHIR connection
type Connection struct {
	BaseURL string
	client  *http.Client
}

// New creates a new connection
func New(baseurl string) *Connection {
	return &Connection{BaseURL: baseurl, client: &http.Client{}}
}

// Query sends a query to the base url
func (c *Connection) Query(q string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v%v", c.BaseURL, q), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
