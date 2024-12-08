package config

import (
	"fmt"

	"github.com/vuonwebadmin/go-admin-packages/config/environment"
)

func Bind[T any](env ...environment.Environment) (T, error) {
	return BindConfig[T]("", env...)
}

func BindConfig[T any](path string, env ...environment.Environment) (T, error) {
	var configPath string
	
	currentEnv := environment.Environment("")
	if len(env) > 0 {
		currentEnv = env[0]
	} else {
		currentEnv = environment.Development
	}

	// TODO: config path
	fmt.Printf("config path: %s\n", configPath)
	fmt.Printf("currentEnv: %s\n", currentEnv)

	return T{}, nil
}
