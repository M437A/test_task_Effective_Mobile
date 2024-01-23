package api

import (
	"errors"
	"fmt"
	"name_service/pkg/api"
)

const (
	nationalityApi    = "https://api.nationalize.io/?name=%s"
	targetNationality = "country"
)

func GetNationality(name string) ([]string, error) {
	url := fmt.Sprintf(nationalityApi, name)
	nationalities, err := api.ApiRequest(url, targetNationality)
	if err != nil {
		return nil, err
	}

	switch res := nationalities.(type) {
	case []interface{}:
		var countries []string
		for _, item := range res {
			countryMap, ok := item.(map[string]interface{})
			if !ok {
				return nil, errors.New("invalid country data in the response")
			}
			countryID, ok := countryMap["country_id"].(string)
			if !ok {
				return nil, errors.New("country_id not found in the response")
			}
			countries = append(countries, countryID)
		}
		return countries, nil
	default:
		return nil, errors.New("country data not found in the response")
	}
}
