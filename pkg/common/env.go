package common

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init_env_file() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}
}
