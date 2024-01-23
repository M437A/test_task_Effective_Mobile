package common

import "database/sql"

func Init() *sql.DB {
	init_env_file()
	init_yaml_file()
	db := createDatabaseConnection()

	return db
}
