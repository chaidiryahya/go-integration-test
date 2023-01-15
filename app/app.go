package app

import (
	// golang package
	"encoding/json"
	"io/ioutil"
	"log"

	// external package
	jsondiff "github.com/nsf/jsondiff"
)

func IntegrationTest() error {
	path := "./scenario/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("[Error] Failed read files from path : %s - IntegrationTest", path)
		return err
	}
	log.Printf("[Integration Test] Found %d scenario file(s)", len(files))

	// var scenarios []Scenario
	for _, file := range files {
		// read file
		fileContent, err := ioutil.ReadFile(path + file.Name())
		if err != nil {
			log.Printf("[Error] Failed read file content : %s - IntegrationTest", file.Name())
			return err
		}

		var scenario Scenario
		err = json.Unmarshal([]byte(fileContent), &scenario)
		if err != nil {
			log.Printf("[Error] Failed unmarshal - IntegrationTest")
			return nil
		}

		// call api
		actualResponse, err := CallAPI(CallAPIParam{
			Host:   scenario.Host,
			Method: scenario.Method,
			Path:   scenario.Path,
		})
		if err != nil {
			log.Printf("%s Failed. Reason : Failed CallAPI [%s] %s%s", scenario.Name, scenario.Method, scenario.Host, scenario.Path)
		}

		expectedResponse, err := json.Marshal(scenario.ExpectedResponse)
		if err != nil {
			log.Printf("%s Failed. Reason : Failed Marshal Expected Response", scenario.Name)
			return nil
		}

		difference, _ := jsondiff.Compare(actualResponse, expectedResponse, &jsondiff.Options{})

		if difference == jsondiff.FullMatch {
			log.Printf("%s => Success", scenario.Name)
		} else {
			log.Printf("%s => Failed. Reason : Expected Response not equal with Current Response", scenario.Name)
		}
	}

	return nil
}
