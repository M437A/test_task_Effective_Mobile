package api

import (
	"name_service/iternal/models"
	rep "name_service/iternal/repository"
)

func Update(userData *models.UserData) (*models.UserData, error) {
	return rep.UpdateUserData(userData)
}
