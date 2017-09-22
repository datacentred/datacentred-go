package datacentred

import (
	"encoding/json"
	"fmt"
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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ProjectResponse struct {
	Projects []Project
}

func ListProjects() []Project {
	data, err := Request("GET", "projects")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res ProjectResponse
	json.Unmarshal(data, &res)
	return res.Projects
}
