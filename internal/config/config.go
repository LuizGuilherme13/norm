package config

import (
	"github.com/joho/godotenv"
)

func GetEnv() (env map[string]string, err error) {
	if err = godotenv.Load("../../../.env"); err != nil {
		return
	}

	env, err = godotenv.Read("../../../.env")
	if err != nil {
		return
	}

	return

}
