package repository

import (
	"name_service/iternal/models"
	"name_service/pkg/common"
)

func SaveUserData(userData *models.UserData) (*models.UserData, error) {
	tx, err := common.DB.Begin()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO user_data (name, surname, patronymic, age, gender, nationality) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, patronymic, age, gender, nationality`

	var savedUserData models.UserData
	err = tx.QueryRow(query, userData.Name, userData.Surname, userData.Patronymic, userData.Age, userData.Gender, userData.Nationality).
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
