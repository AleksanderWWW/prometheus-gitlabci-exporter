package exporter

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var defaultHost string = "gitlab.com"

type gitlabConfig struct {
	Host *string `yaml:"host"`
	ApiToken *string `yaml:"token"`
}

type exporterConfig struct {
	Gitlab gitlabConfig `yaml:"gitlab"`
}

func ParseConfig(path string) (*exporterConfig, error) {
	var config exporterConfig

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	if config.Gitlab.Host == nil {
		config.Gitlab.Host = &defaultHost
	}

	if config.Gitlab.ApiToken == nil {
		token := os.Getenv("GITLAB_API_TOKEN")
		if token == "" {
			return nil, fmt.Errorf("neither config's gitlab.token field nor GITLAB_API_TOKEN env var found")
		}
		config.Gitlab.ApiToken = &token
	}

	return &config, nil
}
