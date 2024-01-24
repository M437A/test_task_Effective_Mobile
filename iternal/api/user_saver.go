package api

import (
	"name_service/iternal/models"
	rep "name_service/iternal/repository"
	api "name_service/pkg/api/privider"
	"sync"
)

func Save(userData *models.UserData) (*models.UserData, error) {
	name := userData.Name

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		age, err := api.GetAge(name)
		if err != nil {
			exceptions(err)
		}
		userData.Age = age
		wg.Done()
	}()

	go func() {
		gender, err := api.GetGender(name)
		if err != nil {
			exceptions(err)
		}
		userData.Gender = gender
		wg.Done()
	}()

	go func() {
		nationality, err := api.GetNationality(name)
		if err != nil {
			exceptions(err)
		}
		userData.Nationality = nationality[0]
		wg.Done()
	}()

	wg.Wait()

	return rep.SaveUserData(userData)
}

func exceptions(err error) error {
	if err != nil {
		return err
	}
	return nil
}
