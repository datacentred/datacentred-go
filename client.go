package datacentred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func PrettyPrintJson(input []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, input, "", "  ")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	} else {
		fmt.Println(string(prettyJSON.Bytes()))
	}
}

func loadCredentialsFromEnv() (string, string) {
	return os.Getenv("DATACENTRED_ACCESS"), os.Getenv("DATACENTRED_SECRET")
}

func Request(verb string, path string) ([]byte, error) {
	client := &http.Client{}

	var url = "https://my.datacentred.io/api/" + path

	req, _ := http.NewRequest(verb, url, nil)
	req.Header.Add("Accept", "application/vnd.datacentred.api+json; version=1")
	access_key, secret_key := loadCredentialsFromEnv()
	req.Header.Add("Authorization", "Token token="+access_key+":"+secret_key)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(url)
	fmt.Println(resp.StatusCode)

	return resp_body, nil
}

func ListUsers() []User {
	data, err := Request("GET", "users")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res UserResponse
	json.Unmarshal(data, &res)
	return res.Users
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

func ListRoles() []Role {
	data, err := Request("GET", "roles")
	if err != nil {
		fmt.Errorf("Request failed: %s", err)
		return nil
	}
	var res RoleResponse
	json.Unmarshal(data, &res)
	return res.Roles
}

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
