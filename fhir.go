package fhir

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	timeout = 3
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
	return &Connection{
		BaseURL: baseurl,
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   time.Duration(timeout*3) * time.Second,
					KeepAlive: time.Duration(timeout*3) * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   time.Duration(timeout) * time.Second,
				ResponseHeaderTimeout: time.Duration(timeout) * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
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
