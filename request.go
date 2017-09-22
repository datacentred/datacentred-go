package datacentred

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiError struct {
	Detail   string
	Field    string
	Resource string
}

type ApiErrors struct {
	Errors []ApiError
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

	switch resp.StatusCode {
	case
		200,
		201,
		204:
		return resp_body, nil
	case
		401:
		return nil, errors.New("Unauthorized: check your credentials")
	case
		404:
		return nil, errors.New("Not found")
	case
		422:
		var apiErrors ApiErrors
		json.Unmarshal(resp_body, &apiErrors)
		return nil, errors.New(apiErrors.Errors[0].Detail)
	}

	return resp_body, nil
}
