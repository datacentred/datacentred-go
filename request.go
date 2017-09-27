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

// ApiError contains details of a single error returned from the API.
type ApiError struct {
	Detail   string
	Field    string
	Resource string
}

type apiErrorsResponse struct {
	Errors []ApiError
}

// ApiVersion is the API microversion this SDK currently targets on server.
const ApiVersion = "1"

// ContentType is the request content type for the API.
const ContentType = "application/json"

// AcceptType is the response type we accept from the server.
const AcceptType = "application/vnd.datacentred.api+json"

// BaseUri is the server's base URI before path suffixes are added.
const BaseUri = "https://my.datacentred.io/api/"

// Configuration is a structure for config data this SDK.
type Configuration struct {
	Client    http.Client
	AccessKey string
	SecretKey string
}

// Config holds config information for this SDK.
// This includes API credentials and a HTTP transport client.
var Config = Configuration{
	Client:    http.Client{},
	AccessKey: os.Getenv("DATACENTRED_ACCESS"),
	SecretKey: os.Getenv("DATACENTRED_SECRET"),
}

// Request makes a single HTTP request against the API server and returns
// the server's response as a byte sequence.
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
	respBody, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case
		200,
		201,
		204:
		return respBody, nil
	case
		401:
		return nil, errors.New("Unauthorized: check your credentials")
	case
		404:
		return nil, errors.New("Not found")
	case
		422:
		var apiErrors apiErrorsResponse
		json.Unmarshal(respBody, &apiErrors)
		fmt.Println(apiErrors.Errors[0].Detail)
		return nil, errors.New(apiErrors.Errors[0].Detail)
	}

	return respBody, nil
}
