package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type MongoDBCollection struct {
	name string `yaml:"name"`
}

type QueueHandlerConfig struct {
	Kafka struct {
		NetworkType                  string `yaml:"networkType"`
		Address                      string `yaml:"address"`
		ProductCreationRequestConfig struct {
			Topic     string `yaml:"topic"`
			Partition int    `yaml:"partition"`
		} `yaml:"productCreationRequest"`
	} `yaml:"kafka"`

	MongoDB struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
	} `yaml:"mongodb"`

	SQL struct {
		NetworkType string `yaml:"networkType"`
		Address     string `yaml:"address"`
		DbName      string `yaml:"dbName"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`

		DriverName string `yaml:"driverName"`
	}
}

func NewQueueHandlerConfig(configPath string) (*QueueHandlerConfig, error) {
	config := &QueueHandlerConfig{}

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
