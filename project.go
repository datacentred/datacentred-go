package datacentred

import (
	"encoding/json"
	"fmt"
	"time"
)

type QuotaSet struct {
	Compute struct {
		Cores    int `json:"cores"`
		Instances int `json:"instances"`
		Ram      int `json:"ram"`
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

type Project struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	QuotaSet QuotaSet `json:"quota_set"`
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

func CreateProject(Params interface{}) Project {
	jsonStr, _ := json.Marshal(Params)
	data, err := Request("POST", "projects", jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res ProjectResponse
	json.Unmarshal(data, &res)
	return res.Project
}

func UpdateProject(Id string, Params interface{}) Project {
	jsonStr, _ := json.Marshal(Params)
	data, err := Request("PUT", "projects/"+Id, jsonStr)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res ProjectResponse
	json.Unmarshal(data, &res)
	return res.Project
}

func DestroyProject(Id string) bool {
	_, err := Request("DELETE", "projects/"+Id, nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	return true
}
