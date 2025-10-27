package exporter

import (
    "os"

    "gopkg.in/yaml.v2"
)


type gitlabConfig struct {
	Host string `yaml:"host"`
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

	return &config, nil
}
