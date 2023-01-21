package app

import (
	// golang package
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	// external package
	jsondiff "github.com/nsf/jsondiff"
)

const (
	AllModule     = "ALL"
	ModuleAPI     = "API"
	ModuleAPIPath = "./test_case/api/"
	ModuleGQL     = "GQL"
	ModuleGQLPath = "./test_case/gql/"
)

func IntegrationTest(Module string) error {

	var paths []string
	switch Module {
	case AllModule:
		paths = append(paths, ModuleAPIPath, ModuleGQLPath)
	case ModuleAPI:
		paths = append(paths, ModuleAPIPath)
	case ModuleGQL:
		paths = append(paths, ModuleGQLPath)
	default:
		return errors.New("Invalid Module Name")
	}

	log.Printf("[Integration Test] Testing %s Module(s)", Module)

	for _, path := range paths {

		var currentModule string
		switch true {
		case strings.Contains(path, "api"):
			currentModule = ModuleAPI
		case strings.Contains(path, "gql"):
			currentModule = ModuleGQL
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Printf("[Error] Failed read files from path : %s - IntegrationTest", path)
			return err
		}

		log.Printf("[Integration Test] Found %d %s test case file(s)", len(files), currentModule)

		for _, file := range files {

			// read file
			fileContent, err := ioutil.ReadFile(path + file.Name())
			if err != nil {
				log.Printf("[Error] Failed read file content : %s - IntegrationTest", file.Name())
				return err
			}

			var testCase TestCase
			err = json.Unmarshal([]byte(fileContent), &testCase)
			if err != nil {
				log.Printf("[Error] Failed unmarshal - IntegrationTest")
				return nil
			}

			headers := []Header{}
			for _, h := range testCase.Header {
				headers = append(headers, Header{
					Key:   h.Key,
					Value: h.Value,
				})
			}

			// call api
			actualResponse, err := CallAPI(CallAPIParam{
				Host:    testCase.Host,
				Method:  testCase.Method,
				Path:    testCase.Path,
				Payload: testCase.Payload,
				Headers: headers,
			})
			if err != nil {
				log.Printf("%s Failed. Reason : Failed CallAPI [%s] %s%s", testCase.Name, testCase.Method, testCase.Host, testCase.Path)
			}

			expectedResponse, err := json.Marshal(testCase.ExpectedResponse)
			if err != nil {
				log.Printf("%s Failed. Reason : Failed Marshal Expected Response", testCase.Name)
				return nil
			}

			difference, _ := jsondiff.Compare(actualResponse, expectedResponse, &jsondiff.Options{})

			if difference == jsondiff.FullMatch {
				log.Printf("[%s] %s => Success", currentModule, testCase.Name)
			} else {
				log.Printf("[%s] %s => Failed. Reason : Expected Response not equal with Current Response", currentModule, testCase.Name)
			}
		}
	}

	return nil
}
