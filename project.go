package datacentred

import (
	"encoding/json"
	"time"
)

// QuotaSet is a collection of resource quota values for a Project.
// It contains Compute, Volume, and Network values as positive integers.
type QuotaSet struct {
	Compute struct {
		Cores     int `json:"cores"`
		Instances int `json:"instances"`
		Ram       int `json:"ram"`
	} `json:"compute"`
	Volume struct {
		Gigabytes int `json:"gigabytes"`
		Snapshots int `json:"snapshots"`
		Volumes   int `json:"volumes"`
	} `json:"volume"`
	Network struct {
		FloatingIp        int `json:"floating_ip"`
		Network           int `json:"network"`
		Port              int `json:"port"`
		Router            int `json:"router"`
		SecurityGroup     int `json:"security_group"`
		SecurityGroupRule int `json:"security_group_rule"`
		Subnet            int `json:"subnet"`
	} `json:"network"`
}

// Project is a cloud organizational unit.
// It is possible to manage a project's name and resource quotas,
// as well as assign/revoke user access.
type Project struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	QuotaSet  QuotaSet  `json:"quota_set"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type projectsResponse struct {
	Projects []Project `json:"projects"`
}

type projectResponse struct {
	Project Project `json:"project"`
}

// Projects is the collection of projects belonging to the currently authenticated user's account.
func Projects() ([]Project, error) {
	data, err := Request("GET", "projects", nil)
	if err != nil {
		return nil, err
	}
	var res projectsResponse
	json.Unmarshal(data, &res)
	return res.Projects, nil
}

// FindProject locates a specific project by its unique ID.
func FindProject(Id string) (*Project, error) {
	data, err := Request("GET", "projects/"+Id, nil)
	if err != nil {
		return nil, err
	}
	var res projectResponse
	json.Unmarshal(data, &res)
	return &res.Project, nil
}

// CreateProject creates a new project on the currently authenticated user's account.
// This may fail if the account is at its limit for projects.
func CreateProject(Params interface{}) (*Project, error) {
	project := map[string]interface{}{
		"project": Params,
	}
	jsonStr, _ := json.Marshal(project)
	data, err := Request("POST", "projects", jsonStr)
	if err != nil {
		return nil, err
	}
	var res projectResponse
	json.Unmarshal(data, &res)
	return &res.Project, nil
}

// Save commits any changes to this project's fields.
func (p Project) Save() (Project, error) {
	project := map[string]interface{}{
		"project": p,
	}
	jsonStr, _ := json.Marshal(project)
	data, err := Request("PUT", "projects/"+p.Id, jsonStr)
	if err != nil {
		return p, err
	}
	var res projectResponse
	json.Unmarshal(data, &res)
	return res.Project, nil
}

// Destroy permanently removes this project.
func (p Project) Destroy() (bool, error) {
	_, err := Request("DELETE", "projects/"+p.Id, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Users is the collection of users assigned to this project.
func (p Project) Users() ([]User, error) {
	data, err := Request("GET", "projects/"+p.Id+"/users", nil)
	if err != nil {
		return nil, err
	}
	var res usersResponse
	json.Unmarshal(data, &res)
	return res.Users, nil
}

// AddUser grants a user access to this project.
func (p Project) AddUser(UserId string) (bool, error) {
	_, err := Request("PUT", "projects/"+p.Id+"/users/"+UserId, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// RemoveUser revokes a user's access to this project.
func (p Project) RemoveUser(UserId string) (bool, error) {
	_, err := Request("DELETE", "projects/"+p.Id+"/users/"+UserId, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
