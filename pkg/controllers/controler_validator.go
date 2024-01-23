package controllers

import (
	"errors"
	"name_service/iternal/models"
)

func ValidateSaveUserDataModel(userData *models.UserData) error {

	if userData.Name == "" ||
		userData.Surname == "" {
		return errors.New("First and Last Name are required")
	}
	return nil
}

func ValidateUpdateUserDataModel(userData *models.UserData) error {
	if userData.Id == 0 ||
		userData.Name == "" ||
		userData.Surname == "" {
		return errors.New("ID, name and surname are required")
	}

	return nil
}

func ValidateGetUserDataModel(userRequest *models.UserDataRequest) error {
	if userRequest.Page == 0 {
		return errors.New("Page is required")
	}
	return nil
}
