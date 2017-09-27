package datacentred

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type UsageBreakdown struct {
	Cost struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Value    float64 `json:"value"`
	}
	Meta  map[string]interface{} `json:"meta"`
	Unit  string                 `json:"unit"`
	Value float64                `json:"value"`
}

type ProjectImagesUsage struct {
	CreatedAt    time.Time        `json:"created_at"`
	DeletedAt    time.Time        `json:"deleted_at"`
	Id           string           `json:"id"`
	LatestSizeGb float64          `json:"latest_size_gb"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	Usage        []UsageBreakdown `json:"usage"`
}

type InstanceHistoryEntry struct {
	Billable   bool      `json:"billable"`
	EventName  string    `json:"event_name"`
	Flavor     string    `json:"flavor"`
	RecordedAt time.Time `json:"recorded_at"`
	Seconds    int       `json:"seconds"`
	State      string    `json:"state"`
	UserId     string    `json:"user_id"`
}

type InstanceFlavor struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	RamMb      int    `json:"ram_mb"`
	RootDiskGb int    `json:"root_disk_gb"`
	VcpusCount int    `json:"vcpus_count"`
}

type ProjectInstancesUsage struct {
	FirstBootedAt time.Time              `json:"first_booted_at"`
	TerminatedAt  time.Time              `json:"terminated_at"`
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	LatestState   string                 `json:"latest_state"`
	History       []InstanceHistoryEntry `json:"history"`
	Tags          []string               `json:"tags"`
	CurrentFlavor InstanceFlavor         `json:"current_flavor"`
	Usage         []UsageBreakdown       `json:"usage"`
}

type IpQuotaChange struct {
	Previous   int       `json:"previous"`
	Quota      int       `json:"quota"`
	RecordedAt time.Time `json:"recorded_at"`
}

type ProjectIpsUsage struct {
	CurrentQuota int              `json:"current_quota"`
	QuotaChanges []IpQuotaChange  `json:"quota_changes"`
	Usage        []UsageBreakdown `json:"usage"`
}

type ProjectLoadBalancersUsage struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	StartedAt    time.Time        `json:"started_at"`
	TerminatedAt time.Time        `json:"terminated_at"`
	Usage        []UsageBreakdown `json:"usage"`
}

type ProjectVolumesUsage struct {
	CreatedAt    time.Time        `json:"created_at"`
	DeletedAt    time.Time        `json:"deleted_at"`
	Id           string           `json:"id"`
	LatestSizeGb int              `json:"latest_size_gb"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	Tags         []string         `json:"tags"`
	Usage        []UsageBreakdown `json:"usage"`
}

type ProjectVpnsUsage struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	StartedAt    time.Time        `json:"started_at"`
	TerminatedAt time.Time        `json:"terminated_at"`
	Usage        []UsageBreakdown `json:"usage"`
}

type ProjectUsage struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Usage struct {
		Images        []ProjectImagesUsage        `json:"images"`
		Instances     []ProjectInstancesUsage     `json:"instances"`
		Ips           ProjectIpsUsage             `json:"ips"`
		LoadBalancers []ProjectLoadBalancersUsage `json:"load_balancers"`
		ObjectStorage struct {
			Usage []UsageBreakdown `json:"usage"`
		} `json:"object_storage"`
		Volumes []ProjectVolumesUsage `json:"volumes"`
		Vpns    []ProjectVpnsUsage    `json:"vpns"`
	} `json:"usage"`
}

type Usage struct {
	LastUpdatedAt time.Time      `json:"last_updated_at"`
	Projects      []ProjectUsage `json:"projects"`
}

func ShowUsage(year int, month int) Usage {
	data, err := Request("GET", "usage/"+strconv.Itoa(year)+"/"+strconv.Itoa(month), nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res Usage
	json.Unmarshal(data, &res)
	return res
}
