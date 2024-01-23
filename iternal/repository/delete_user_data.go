package repository

import "name_service/pkg/common"

func DeleteByID(id int64) error {
	tx, err := common.DB.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM user_data WHERE id = $1`
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
