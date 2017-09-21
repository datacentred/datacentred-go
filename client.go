package main

import (
  "bytes"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

func Request(verb string, path string) {
  client := &http.Client{}

  req, _ := http.NewRequest(verb, path, nil)
  req.Header.Add("Accept", "application/vnd.datacentred.api+json; version=1")
  req.Header.Add("Authorization", "Token token=foo:bar")
  req.Header.Add("Content-Type", "application/json")

  resp, _ := client.Do(req)
  s, _    := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var prettyJSON bytes.Buffer
  error := json.Indent(&prettyJSON, s, "", "  ")
  if error != nil {
      fmt.Println("JSON parse error: ", error)
  }

  fmt.Println((string(prettyJSON.Bytes())))
  return
}

func main() {
  Request("GET", "https://my.datacentred.io/api")
}