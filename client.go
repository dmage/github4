package github4

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var DefaultEndpoint = "https://api.github.com/graphql"

type Client struct {
	Endpoint string
	Client   *http.Client
}

func (c *Client) Do(query string, variables map[string]interface{}, response interface{}) error {
	q := struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables,omitempty"`
	}{
		Query:     query,
		Variables: variables,
	}
	buf, err := json.Marshal(q)
	if err != nil {
		return err
	}

	endpoint := c.Endpoint
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}

	httpClient := c.Client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Post(endpoint, "", bytes.NewReader(buf))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var v struct {
		Data   json.RawMessage
		Errors []struct {
			Message string
		}
	}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		return err
	}

	if len(v.Errors) > 0 {
		return fmt.Errorf("%s", v.Errors[0].Message)
	}

	return json.Unmarshal(v.Data, response)
}
