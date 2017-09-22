package datacentred

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func loadCredentialsFromEnv() (string, string) {
	return os.Getenv("DATACENTRED_ACCESS"), os.Getenv("DATACENTRED_SECRET")
}
