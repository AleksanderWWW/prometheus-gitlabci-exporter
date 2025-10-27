package exporter

import (
    "os"

    "gopkg.in/yaml.v2"
)


var defaultHost string = "gitlab.com"
var defaultPort int32  = 443


type gitlabConfig struct {
	Host *string `yaml:"host"`
	Port *int32 `yaml:"port"`
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

	if config.Gitlab.Port == nil {
		config.Gitlab.Port = &defaultPort
	}

	return &config, nil
}
