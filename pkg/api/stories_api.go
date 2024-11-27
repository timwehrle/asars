package api

import (
	"fmt"
	"net/http"
)

type Story struct {
	GID       string `json:"gid"`
	CreatedBy User   `json:"created_by"`
	Text      string `json:"text"`
}

func (c *Client) GetStories(workspaceGID, taskGID string) ([]Story, error) {
	resp, err := c.makeRequest(http.MethodGet, "/tasks/"+taskGID+"/stories", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []Story `json:"data"`
	}

	if err := handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to handle response: %w", err)
	}

	return result.Data, nil
}
