package api

import (
	"errors"
	"fmt"
	"name_service/pkg/api"
)

const (
	genderAPI    = "https://api.genderize.io/?name=%s"
	targetGender = "gender"
)

func GetGender(name string) (string, error) {
	url := fmt.Sprintf(genderAPI, name)
	gender, err := api.ApiRequest(url, targetGender)
	if err != nil {
		return "", err
	}

	switch res := gender.(type) {
	case string:
		return res, nil
	default:
		return "", errors.New("gender not found in the response")
	}
}
