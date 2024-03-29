package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	WebInterface struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"web_interface"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	cwd, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current working directory: %v\n", err)
	}

	execPwd := filepath.Dir(cwd)

	config, err := parseConfig(execPwd + "/configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return config, nil
}

func parseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML data: %v", err)
	}

	return config, nil
}
