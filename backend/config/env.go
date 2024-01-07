package config

import (
	"os"
	"strings"
)

const dotEnvFilePath = "./local.env"

func SetAllEnvs() {
	env, err := os.ReadFile(dotEnvFilePath)
	if err != nil {
		panic("Erro ao carregar .env")
	}
	parsedEnv := string(env)
	splitedEnv := strings.Split(parsedEnv, "\n")
	for i := 0; i < len(splitedEnv); i++ {
		splitedItem := strings.Split(splitedEnv[i], "=")
		os.Setenv(splitedItem[0], splitedItem[1])
	}
}
