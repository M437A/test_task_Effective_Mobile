package repository

import (
	"name_service/iternal/models"
	"name_service/pkg/common"
)

func UpdateUserData(userData *models.UserData) (*models.UserData, error) {
	tx, err := common.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `UPDATE user_data 
		SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 
		WHERE id = $7
		RETURNING id, name, surname, patronymic, age, gender, nationality`

	var savedUserData models.UserData
	err = tx.QueryRow(query, userData.Name, userData.Surname, userData.Patronymic, userData.Age, userData.Gender, userData.Nationality, userData.Id).
		Scan(&savedUserData.Id, &savedUserData.Name, &savedUserData.Surname, &savedUserData.Patronymic, &savedUserData.Age, &savedUserData.Gender, &savedUserData.Nationality)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &savedUserData, nil
}
