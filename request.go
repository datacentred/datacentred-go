package datacentred

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
