package datacentred

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type UsageBreakdown struct {
	Cost struct {
		Currency string
		Rate     float64
		Value    float64
	}
	Meta  map[string]interface{}
	Unit  string
	Value float64
}

type ProjectImagesUsage struct {
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Id           string
	LatestSizeGb float64 `json:"latest_size_gb"`
	Name         string
	Owner        string
	Usage        []UsageBreakdown
}

type InstanceHistoryEntry struct {
	Billable   bool
	EventName  string `json:"event_name"`
	Flavor     string
	RecordedAt time.Time `json:"recorded_at"`
	Seconds    int
	State      string
	UserId     string `json:"user_id"`
}

type InstanceFlavor struct {
	Id         string
	Name       string
	RamMb      int `json:"ram_mb"`
	RootDiskGb int `json:"root_disk_gb"`
	VcpusCount int `json:"vcpus_count"`
}

type ProjectInstancesUsage struct {
	FirstBootedAt time.Time `json:"first_booted_at"`
	TerminatedAt  time.Time `json:"terminated_at"`
	Id            string
	Name          string
	LatestState   string `json:"latest_state"`
	History       []InstanceHistoryEntry
	Tags          []string
	CurrentFlavor InstanceFlavor `json:"current_flavor"`
	Usage         []UsageBreakdown
}

type IpQuotaChange struct {
	Previous   int
	Quota      int
	RecordedAt time.Time `json:"recorded_at"`
}

type ProjectIpsUsage struct {
	CurrentQuota int             `json:"current_quota"`
	QuotaChanges []IpQuotaChange `json:"quota_changes"`
	Usage        []UsageBreakdown
}

type ProjectLoadBalancersUsage struct {
	Id           string
	Name         string
	Owner        string
	StartedAt    time.Time `json:"started_at"`
	TerminatedAt time.Time `json:"terminated_at"`
	Usage        []UsageBreakdown
}

type ProjectVolumesUsage struct {
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Id           string
	LatestSizeGb int `json:"latest_size_gb"`
	Name         string
	Owner        string
	Tags         []string
	Usage        []UsageBreakdown
}

type ProjectVpnsUsage struct {
	Id           string
	Name         string
	Owner        string
	StartedAt    time.Time `json:"started_at"`
	TerminatedAt time.Time `json:"terminated_at"`
	Usage        []UsageBreakdown
}

type ProjectUsage struct {
	Id    string
	Name  string
	Usage struct {
		Images        []ProjectImagesUsage
		Instances     []ProjectInstancesUsage
		Ips           ProjectIpsUsage
		LoadBalancers []ProjectLoadBalancersUsage `json:"load_balancers"`
		ObjectStorage struct {
			Usage []UsageBreakdown
		} `json:"object_storage"`
		Volumes []ProjectVolumesUsage
		Vpns    []ProjectVpnsUsage
	}
}

type Usage struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Projects      []ProjectUsage
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
