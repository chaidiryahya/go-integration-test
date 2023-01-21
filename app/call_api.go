package app

import (
	// golang package
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// CallAPI will call api
//
// It returns slice of byte and nil error when successfull
// Otherwise nil byte and non nil error will be returned
func CallAPI(param CallAPIParam) ([]byte, error) {
	url := param.Host + param.Path

	// Add payload
	var payload io.Reader
	if param.Payload != "" {
		payload = strings.NewReader(param.Payload)
	}

	req, err := http.NewRequest(param.Method, url, payload)
	if err != nil {
		log.Printf("[Error] Failed http.NewRequest() - CallAPI")
		return nil, err
	}

	// Add request header
	for _, header := range param.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("[Error] Failed client.Do() - CallAPI")
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("[Error] Failed ioutil.ReadAll() - CallAPI")
		return nil, err
	}

	return body, nil
}
