package config

import (
	"encoding/json"
	"os"
)

type ItemConfig struct {
	Name      string  `json:"name"`
	URL       string  `json:"url"`
	Selector  string  `json:"selector"`
	Threshold float64 `json:"threshold"`
}

func LoadConfig(path string) ([]ItemConfig, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result []ItemConfig
	if err = json.Unmarshal(fileBytes, &result); err != nil {
		return nil, err
	}
	return result, nil
}
