package datacentred

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

type Configuration struct {
	Client    http.Client
	AccessKey string
	SecretKey string
}

var Config = Configuration{
	Client:    http.Client{},
	AccessKey: os.Getenv("DATACENTRED_ACCESS"),
	SecretKey: os.Getenv("DATACENTRED_SECRET"),
}

func Request(verb string, path string, body []byte) ([]byte, error) {
	var url = BaseUri + path

	req, _ := http.NewRequest(verb, url, bytes.NewBuffer(body))
	req.Header.Add("Accept", AcceptType+"; version="+ApiVersion)
	req.Header.Add("Authorization", "Token token="+Config.AccessKey+":"+Config.SecretKey)
	req.Header.Add("Content-Type", ContentType)

	resp, err := Config.Client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	resp_body, _ := ioutil.ReadAll(resp.Body)

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
