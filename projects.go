package redmineclientgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// OBS: We use this struct since redmine API is a bit weird like that
// ref: https://www.redmine.org/projects/redmine/wiki/Rest_Projects
type ProjectCreateRequest struct {
	Project *Project `json:"project"`
}

func (c *Client) CreateProject(project Project) (*Project, error) {
	projectRequest := &ProjectCreateRequest{
		Project: &project,
	}

	projectRequestJson, err := json.Marshal(projectRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/projects.xml", c.HostURL), strings.NewReader(string(projectRequestJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	projectCreated := ProjectCreateRequest{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return projectCreated.Project, nil
}
