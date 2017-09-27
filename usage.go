package datacentred

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// UsageBreakdown show the cost and unit-usage of a resource.
// This structure is common to all resource types.
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

// ProjectImagesUsage contains usage information for a single cloud image.
type ProjectImagesUsage struct {
	CreatedAt    time.Time        `json:"created_at"`
	DeletedAt    time.Time        `json:"deleted_at"`
	Id           string           `json:"id"`
	LatestSizeGb float64          `json:"latest_size_gb"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	Usage        []UsageBreakdown `json:"usage"`
}

// InstanceHistoryEntry contains information about a specific event in
// a cloud instance's runtime history.
type InstanceHistoryEntry struct {
	Billable   bool      `json:"billable"`
	EventName  string    `json:"event_name"`
	Flavor     string    `json:"flavor"`
	RecordedAt time.Time `json:"recorded_at"`
	Seconds    int       `json:"seconds"`
	State      string    `json:"state"`
	UserId     string    `json:"user_id"`
}

// InstanceFlavor contains information about a cloud instance's assigned flavor.
type InstanceFlavor struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	RamMb      int    `json:"ram_mb"`
	RootDiskGb int    `json:"root_disk_gb"`
	VcpusCount int    `json:"vcpus_count"`
}

// ProjectInstancesUsage contains usage information about a single cloud instance.
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

// IpQuotaChange contains information about an IP quota change event.
type IpQuotaChange struct {
	Previous   int       `json:"previous"`
	Quota      int       `json:"quota"`
	RecordedAt time.Time `json:"recorded_at"`
}

// ProjectIpsUsage contains usage information about IPs.
type ProjectIpsUsage struct {
	CurrentQuota int              `json:"current_quota"`
	QuotaChanges []IpQuotaChange  `json:"quota_changes"`
	Usage        []UsageBreakdown `json:"usage"`
}

// ProjectLoadBalancersUsage contains usage information about a single cloud load balancer.
type ProjectLoadBalancersUsage struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	StartedAt    time.Time        `json:"started_at"`
	TerminatedAt time.Time        `json:"terminated_at"`
	Usage        []UsageBreakdown `json:"usage"`
}

// ProjectVolumesUsage contains usage information about a single cloud volume.
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

// ProjectVpnsUsage contains usage information about a single cloud VPN.
type ProjectVpnsUsage struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Owner        string           `json:"owner"`
	StartedAt    time.Time        `json:"started_at"`
	TerminatedAt time.Time        `json:"terminated_at"`
	Usage        []UsageBreakdown `json:"usage"`
}

// ProjectUsage contains usage information about a single cloud project.
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

// Usage contains usage information about all cloud projects belonging
// to the currently authenticated user's account.
type Usage struct {
	LastUpdatedAt time.Time      `json:"last_updated_at"`
	Projects      []ProjectUsage `json:"projects"`
}

// ShowUsage retrieves resource usage data for a given month belonging to the currently authenticated user's account.
func ShowUsage(year int, month int) Usage {
	data, err := Request("GET", "usage/"+strconv.Itoa(year)+"/"+strconv.Itoa(month), nil)
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	}
	var res Usage
	json.Unmarshal(data, &res)
	return res
}
