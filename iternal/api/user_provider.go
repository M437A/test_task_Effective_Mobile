package api

import (
	"name_service/iternal/models"
	rep "name_service/iternal/repository"
)

func GetFilteredUsers(requestData *models.UserDataRequest) ([]models.UserData, error) {
	filter := requestData.UserFilter
	users, err := rep.GetFilteredUsers(&filter, requestData.Page)
	if err != nil {
		return nil, err
	}
	return users, nil
}
