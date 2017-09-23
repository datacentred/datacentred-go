package datacentred

import (
	"bytes"
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

const ApiVersion = "1"
const ContentType = "application/json"
const AcceptType = "application/vnd.datacentred.api+json"
const BaseUri = "https://my.datacentred.io/api/"

func Request(verb string, path string, body []byte) ([]byte, error) {
	client := http.Client{}

	var url = BaseUri + path

	fmt.Println(bytes.NewBuffer(body))

	req, _ := http.NewRequest(verb, url, bytes.NewBuffer(body))
	req.Header.Add("Accept", AcceptType+"; version="+ApiVersion)
	access_key, secret_key := loadCredentialsFromEnv()
	req.Header.Add("Authorization", "Token token="+access_key+":"+secret_key)
	req.Header.Add("Content-Type", ContentType)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	resp_body, _ := ioutil.ReadAll(resp.Body)

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
		fmt.Println(apiErrors.Errors[0].Detail)
		return nil, errors.New(apiErrors.Errors[0].Detail)
	}

	return resp_body, nil
}
