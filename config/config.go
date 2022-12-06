package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config interface {
	GetString(key string, default_value string) string
	GetInt(key string, default_value string) (int, error)
}

type configImpl struct {
}

func (config *configImpl) GetString(key string, default_value string) string {
	key = strings.TrimSpace(key)
	envVal := os.Getenv(key)
	if len(envVal) > 0 {
		return envVal
	}
	return default_value
}

func (config *configImpl) GetInt(key string, default_value string) (int, error) {
	intVar, err := strconv.Atoi(config.GetString(key, default_value))
	return intVar, err
}

func New(filenames ...string) (Config, error) {
	err := godotenv.Load(filenames...)
	return &configImpl{}, err
}
