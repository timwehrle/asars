package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type Task struct {
	GID       string `json:"gid"`
	Name      string `json:"name"`
	DueOn     string `json:"due_on"`
	CreatedBy User   `json:"created_by"`
	HtmlNotes string `json:"html_notes"`
	Notes     string `json:"notes"`
	Assignee  User   `json:"assignee"`
	Tags      []Tag  `json:"tags"`
	PermaLink string `json:"permalink_url"`
}

func (c *Client) GetTasks(workspaceGID string) ([]Task, error) {
	endpoint := fmt.Sprintf("/tasks?workspace=%s", url.QueryEscape(workspaceGID))
	endpoint += "&opt_fields=name,due_on,completed"
	endpoint += "&completed_since=now"
	endpoint += "&assignee=me"

	resp, err := c.makeRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data []Task `json:"data"`
	}

	if err := handleResponse(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to handle response: %w", err)
	}

	return result.Data, nil
}
