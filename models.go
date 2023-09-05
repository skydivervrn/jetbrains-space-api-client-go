package jetbrains_space_api_client_go

import "net/http"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// Project struct
type Project struct {
	ID  string `json:"id"`
	Key struct {
		Key string `json:"key"`
	} `json:"key"`
	Name                     string `json:"name"`
	Private                  bool   `json:"private"`
	Description              string `json:"description"`
	Icon                     string `json:"icon"`
	LatestRepositoryActivity string `json:"latestRepositoryActivity"`
	CreatedAt                struct {
		Iso       string `json:"iso"`
		Timestamp int64  `json:"timestamp"`
	} `json:"createdAt"`
	Archived bool `json:"archived"`
}

// CreateProjectData struct
type CreateProjectData struct {
	Name string `json:"name"`
	Key  struct {
		Key string `json:"key"`
	} `json:"key"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// UpdateProjectData struct
type UpdateProjectData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// ProjectRepos struct
type ProjectRepos struct {
	Repos []struct {
		ID                        string `json:"id"`
		Name                      string `json:"name"`
		Description               string `json:"description"`
		LatestActivity            string `json:"latestActivity"`
		ProxyPushNotification     string `json:"proxyPushNotification"`
		ProxyPushNotificationBody string `json:"proxyPushNotificationBody"`
		State                     string `json:"state"`
		InitProgress              string `json:"initProgress"`
		ReadmeName                string `json:"readmeName"`
		MonthlyActivity           string `json:"monthlyActivity"`
		DefaultBranch             struct {
			Head string `json:"head"`
			Ref  string `json:"ref"`
		} `json:"defaultBranch"`
	} `json:"repos"`
}

// Repository struct
type Repository struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	LatestActivity            string `json:"latestActivity"`
	ProxyPushNotification     string `json:"proxyPushNotification"`
	ProxyPushNotificationBody string `json:"proxyPushNotificationBody"`
	State                     string `json:"state"`
	InitProgress              string `json:"initProgress"`
	ReadmeName                string `json:"readmeName"`
	MonthlyActivity           string `json:"monthlyActivity"`
	DefaultBranch             struct {
		Head string `json:"head"`
		Ref  string `json:"ref"`
	} `json:"defaultBranch"`
}

// CreateRepositoryData struct
type CreateRepositoryData struct {
	Description   string `json:"description"`
	DefaultBranch string `json:"defaultBranch"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	ProjectId     string `json:"project_id"`
}
