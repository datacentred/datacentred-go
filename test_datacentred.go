package main

import (
	"datacentred"
	"fmt"
)

func main() {
	data, err := datacentred.Request("GET", "users")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
	} else {
		datacentred.PrettyPrintJson(data)
	}
}
