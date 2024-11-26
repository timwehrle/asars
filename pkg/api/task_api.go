package api

import (
	"fmt"
	"net/http"
)

func (c *Client) GetTask(workspaceGID, taskGID string) (*Task, error) {
	resp, err := c.makeRequest(http.MethodGet, "/tasks/"+taskGID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data Task `json:"data"`
	}

	if err := handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to handle response: %w", err)
	}

	return &result.Data, nil
}
