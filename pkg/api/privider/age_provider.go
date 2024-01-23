package api

import (
	"errors"
	"fmt"
	"name_service/pkg/api"
)

const (
	ageAPI    = "https://api.agify.io/?name=%s"
	targetAge = "age"
)

func GetAge(name string) (int64, error) {
	url := fmt.Sprintf(ageAPI, name)
	age, err := api.ApiRequest(url, targetAge)
	if err != nil {
		return 0, err
	}

	switch res := age.(type) {
	case float64:
		return int64(res), nil
	default:
		return 0, errors.New("age not found in the response")
	}
}
