package api

import (
	"name_service/iternal/models"
	rep "name_service/iternal/repository"
	api "name_service/pkg/api/privider"
)

func Save(userData *models.UserData) (*models.UserData, error) {
	name := userData.Name

	age, err := api.GetAge(name)
	exceptions(err)

	gender, err := api.GetGender(name)
	exceptions(err)

	nationality, err := api.GetNationality(name)
	exceptions(err)

	userData.Age = age
	userData.Gender = gender
	userData.Nationality = nationality[0]

	return rep.SaveUserData(userData)
}

func exceptions(err error) error {
	if err != nil {
		return err
	}
	return nil
}
