package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ApiRequest(url string, target string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	value, ok := result[target]
	if !ok {
		return nil, errors.New(target + " not found in the response")
	}

	return value, nil
}
