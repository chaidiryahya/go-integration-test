package app

import (
	// golang package
	"io/ioutil"
	"log"
	"net/http"
)

// CallAPI will call api
//
// It returns slice of byte and nil error when successfull
// Otherwise nil byte and non nil error will be returned
func CallAPI(param CallAPIParam) ([]byte, error) {
	url := param.Host + param.Path

	client := &http.Client{}
	req, err := http.NewRequest(param.Method, url, nil)
	if err != nil {
		log.Printf("[Error] Failed http.NewRequest() - CallAPI")
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("[Error] Failed client.Do() - CallAPI")
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("[Error] Failed ioutil.ReadAll() - CallAPI")
		return nil, err
	}

	return body, nil
}
