package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type DataIngestionApiConfig struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	Kafka struct {
		NetworkType                  string `yaml:"networkType"`
		Address                      string `yaml:"address"`
		ProductCreationRequestConfig struct {
			Topic     string `yaml:"topic"`
			Partition int    `yaml:"partition"`
		} `yaml:"productCreationRequest"`
	} `yaml:"kafka"`
}

func NewDataIngestionApiConfiguration(configPath string) (*DataIngestionApiConfig, error) {
	config := &DataIngestionApiConfig{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
