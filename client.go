package datacentred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func LoadCredentialsFromEnv() (string,string) {
	return os.Getenv("DATACENTRED_ACCESS"), os.Getenv("DATACENTRED_SECRET")
}

func Request(verb string, path string) ([]byte, error) {
	client := &http.Client{}

	var url = "https://my.datacentred.io/api/" + path

	req, _ := http.NewRequest(verb, url, nil)
	req.Header.Add("Accept", "application/vnd.datacentred.api+json; version=1")
	access_key, secret_key := LoadCredentialsFromEnv()
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

	return resp_body, nil
}
