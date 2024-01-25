package api

import (
	"fmt"
	"name_service/iternal/models"
	rep "name_service/iternal/repository"
	api "name_service/pkg/api/privider"
	"sync"
)

func Save(userData *models.UserData) (*models.UserData, error) {
	name := userData.Name

	var wg sync.WaitGroup
	wg.Add(3)

	var errs []error

	go func() {
		defer wg.Done()
		if age, err := api.GetAge(name); err != nil {
			errs = append(errs, err)
		} else {
			userData.Age = age
		}
	}()

	go func() {
		defer wg.Done()
		if gender, err := api.GetGender(name); err != nil {
			errs = append(errs, err)
		} else {
			userData.Gender = gender
		}
	}()

	go func() {
		defer wg.Done()
		if nationality, err := api.GetNationality(name); err != nil {
			errs = append(errs, err)
		} else {
			userData.Nationality = nationality[0]
		}
	}()

	wg.Wait()

	if len(errs) > 0 {
		return nil, fmt.Errorf("Encountered errors: %v", errs)
	}

	return rep.SaveUserData(userData)
}
