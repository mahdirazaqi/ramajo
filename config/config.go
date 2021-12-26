package config

import (
	"encoding/json"
	"os"
)

type (
	config struct {
		Port         string `json:"port"`
		TemplatePath string `json:"template_path"`
		Mongo        mongo  `json:"mongo"`
	}

	mongo struct {
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

func Load(path string) (*config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := new(config)
	if err := json.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
