package datacentred

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ShowUsage(year int, month int) map[string]interface{} {
	data, err := Request("GET", "usage/"+strconv.Itoa(year)+"/"+strconv.Itoa(month))
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res map[string]interface{}
	json.Unmarshal(data, &res)
	return res
}
