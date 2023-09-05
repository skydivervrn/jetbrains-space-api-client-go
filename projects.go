package jetbrains_space_api_client_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetProject - Returns project by id
func (c *Client) GetProject(id string) (Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/http/projects/id:%s", c.HostURL, id), nil)
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

// getProjectRepos
func (c *Client) getProjectRepos(projectId string) (ProjectRepos, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/id:%s?$fields=repos", c.HostURL, baseApiEndpoint, projectId), nil)
	if err != nil {
		return ProjectRepos{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return ProjectRepos{}, err
	}

	project := ProjectRepos{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return ProjectRepos{}, err
	}

	return project, nil
}

// CreateProject - Creates project with given name
func (c *Client) CreateProject(createProjectData CreateProjectData) (Project, error) {
	createProjectData.Key.Key = strings.ToUpper(strings.ReplaceAll(createProjectData.Name, " ", "-"))
	bytesData, _ := json.Marshal(createProjectData)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, baseApiEndpoint), bytes.NewBuffer(bytesData))
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

// UpdateProject - Creates project with given name
func (c *Client) UpdateProject(updateProjectData UpdateProjectData) (Project, error) {
	bytesData, _ := json.Marshal(updateProjectData)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s%s/id:%s", c.HostURL, baseApiEndpoint, updateProjectData.Id), bytes.NewBuffer(bytesData))
	if err != nil {
		return Project{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Project{}, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return Project{}, err
	}

	return project, nil
}

// DeleteProject - Creates project with given name
func (c *Client) DeleteProject(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/id:%s", c.HostURL, baseApiEndpoint, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
