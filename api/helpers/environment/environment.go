package environment

import (
	"errors"
	"fmt"
	"slices"

	"github.com/gofor-little/env"
)

type Env struct {
	ApiAddr string
	Mode    string // dev, prod
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

	// API_ADDR
	e.ApiAddr, err = env.MustGet("API_ADDR")
	if err != nil {
		return err
	}

	// MODE
	e.Mode = env.Get("MODE", "dev")
	if !slices.Contains(modes, e.Mode) {
		return errors.New("MODE environment variable must be either dev or prod")
	}

	environment = e
	environmentSetup = true

	return nil
}

func Get(key string) (string, error) {
	if !environmentSetup {
		return "", errors.New("environment not setup")
	}

	switch key {
	case "API_ADDR":
		return environment.ApiAddr, nil
	case "MODE":
		return environment.Mode, nil
	default:
		return "", fmt.Errorf("key <%s> not found", key)
	}
}
