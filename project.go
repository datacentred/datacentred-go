package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type Project struct {
	Id       string
	Name     string
	QuotaSet struct {
		Compute struct {
			Cores    int
			Instance int
			Ram      int
		}
		Volume struct {
			Gigabytes int
			Snapshots int
			Volumes   int
		}
		Network struct {
			FloatingIp        int `json:"floating_ip"`
			Network           int
			Port              int
			Router            int
			SecurityGroup     int `json:"security_group"`
			SecurityGroupRule int `json:"security_group_rule"`
			Subnet            int
		}
	} `json:"quota_set"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectsResponse struct {
	Projects []Project
}

type ProjectResponse struct {
	Project Project
}

func ListProjects() []Project {
	data, err := Request("GET", "projects", nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res ProjectsResponse
	json.Unmarshal(data, &res)
	return res.Projects
}

func FindProject(Id string) Project {
	data, err := Request("GET", "projects/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res ProjectResponse
	json.Unmarshal(data, &res)
	return res.Project
}
