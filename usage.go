package datacentred

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type UsageBreakdown struct {
	Cost struct {
		Currency string
		Rate     float32
		Value    float32
	}
	Meta  map[string]interface{}
	Unit  string
	Value float32
}

type ProjectImagesUsage struct {
	CreatedAt    string `json:"created_at"`
	DeletedAt    string `json:"deleted_at"`
	Id           string
	LatestSizeGb float32 `json:"latest_size_gb"`
	Name         string
	Owner        string
	Usage        []UsageBreakdown
}

type ProjectInstancesUsage struct {
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
	Id        string
	Name      string
}

type ProjectIpsUsage struct {
}

type ProjectLoadBalancersUsage struct {
}

type ProjectVolumesUsage struct {
}

type ProjectVpnsUsage struct {
}

type ProjectUsage struct {
	Id    string
	Name  string
	Usage struct {
		Images []ProjectImagesUsage
		// Instance []ProjectInstancesUsage
		// Ips []ProjectIpsUsage
		// LoadBalancers []ProjectLoadBalancersUsage `json:"load_balancers"`
		// ObjectStorage struct {
		//   Usage []UsageBreakdown
		// } `json:"object_storage"`
		// Volumes []ProjectVolumesUsage
		// Vpns []ProjectVpnsUsage
	}
}

type UsageResponse struct {
	LastUpdatedAt string `json:"last_updated_at"`
	Projects      []ProjectUsage
}

func ShowUsage(year int, month int) UsageResponse {
	data, err := Request("GET", "usage/"+strconv.Itoa(year)+"/"+strconv.Itoa(month))
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res UsageResponse
	json.Unmarshal(data, &res)
	return res
}
