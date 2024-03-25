package environment

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/gofor-little/env"
)

type Env struct {
	ApiPort      string
	Mode         string // dev, prod
	RedisHost    string
	RedisPort    string
	RedisDbIndex string
}

var environment *Env
var environmentSetup = false

var modes = []string{"dev", "prod"}

func SetupEnv() error {
	var err error

	e := &Env{}

	if err = env.Load(".env"); err != nil {
		return err
	}

	// MODE
	e.Mode = env.Get("MODE", "dev")
	if !slices.Contains(modes, e.Mode) {
		return fmt.Errorf("MODE environment variable must be either dev or prod. Got: %s", e.Mode)
	}

	// API_PORT
	if e.ApiPort, err = env.MustGet("API_PORT"); err != nil {
		return err
	}

	// REDIS
	if e.RedisHost, err = env.MustGet("REDIS_HOST"); err != nil {
		return err
	}
	e.RedisPort = env.Get("REDIS_PORT", "6379")
	e.RedisDbIndex = env.Get("REDIS_DB_INDEX", "0")

	environment = e
	environmentSetup = true

	return nil
}

func Get(key string) (string, error) {
	if !environmentSetup {
		return "", errors.New("environment not setup")
	}

	switch key {
	case "API_PORT":
		return environment.ApiPort, nil
	case "MODE":
		return environment.Mode, nil
	case "REDIS_ADDR":
		return fmt.Sprintf("%s:%s", environment.RedisHost, environment.RedisPort), nil
	case "REDIS_DB_INDEX":
		return environment.RedisDbIndex, nil
	default:
		return "", fmt.Errorf("key <%s> not found", key)
	}
}

func GetInt(key string) (int, error) {
	if !environmentSetup {
		return 0, errors.New("environment not setup")
	}

	strValue, err := Get(key)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, fmt.Errorf("key <%s> is not a valid integer", key)
	}

	return value, nil
}
