package repository

import (
	"database/sql"
	"fmt"
	"name_service/iternal/models"
	"name_service/pkg/common"
	"reflect"
	"strconv"
)

const (
	PageSize = 10
)

func GetFilteredUsers(filter *models.UserFilter, page int) ([]models.UserData, error) {
	tx, err := common.DB.Begin()
	if err != nil {
		return nil, err
	}

	query, args := buildSQLRequest(filter, page)

	rows, err := tx.Query(query, args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	users, err := buildUsersResult(rows)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func buildSQLRequest(filter *models.UserFilter, page int) (string, []interface{}) {
	query := "SELECT * FROM user_data WHERE TRUE"
	args := make([]interface{}, 0)

	filterValue := reflect.ValueOf(filter).Elem()
	for i := 0; i < filterValue.NumField(); i++ {
		fieldValue := filterValue.Field(i)
		fieldTag := filterValue.Type().Field(i).Tag.Get("json")

		if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
			continue
		}

		if fieldValue.Kind() == reflect.Int64 && fieldValue.Int() == 0 {
			continue
		}

		fieldValStr := fmt.Sprintf("%v", fieldValue.Interface())
		query += " AND CAST(" + fieldTag + " AS TEXT) LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+fieldValStr+"%")
	}

	query += " LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, PageSize, (page-1)*PageSize)

	return query, args
}

func buildUsersResult(rows *sql.Rows) ([]models.UserData, error) {
	var users []models.UserData

	for rows.Next() {
		var user models.UserData
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
