package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// tcClient "github.com/petemoore/taskcluster-client.go/lib"
)

var (
	err error
)

// APIDefinition represents the definition of a REST API, comprising of the URL to the defintion
// of the API in json format, together with a URL to a json schema to validate the definition
type APIDefinition struct {
	URL       string `json:"url"`
	SchemaURL string `json:"schema"`
}

//go:generate
func main() {
	var apis []APIDefinition
	var bytes []byte
	bytes, err = ioutil.ReadFile("apis.json")
	exitOnFail()
	err = json.Unmarshal(bytes, &apis)
	exitOnFail()
	for _, api := range apis {
		fmt.Printf("Downloading API endpoint from %v...\n", api.URL)
		var resp *http.Response
		resp, err = http.Get(api.URL)
		exitOnFail()
		defer resp.Body.Close()
		fmt.Printf("Schema URL is %v\n", api.SchemaURL)
		loadJson(resp.Body)
	}
	fmt.Println("All done.")
}

func exitOnFail() {
	if err != nil {
		log.Fatal(err)
	}
}
