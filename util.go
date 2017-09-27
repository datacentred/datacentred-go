package datacentred

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dnaeon/go-vcr/recorder"
	"net/http"
)

func prettyPrintJson(input []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, input, "", "  ")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	} else {
		fmt.Println(string(prettyJSON.Bytes()))
	}
}

func initRecorder(Name string) *recorder.Recorder {
	r, err := recorder.New(Name)
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	}
	defer r.Stop()

	Config.Client = http.Client{
		Transport: r,
	}
	return r
}
