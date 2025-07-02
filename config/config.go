package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"access_manager/internal/common/alert_sender"
	"access_manager/internal/common/postgres_client"
	"access_manager/internal/common/token_generator"
)

type Config struct {
	Postgres       postgres_client.Config `yaml:"postgres"`
	TokenGenerator token_generator.Config `yaml:"token_generator"`
	AlertSender    alert_sender.Config    `yaml:"alert_sender"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig(configPath string) (*Config, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to find configuration file")
	}

	cfg := &Config{}

	if err := yaml.Unmarshal([]byte(string(content)), cfg); err != nil {
		return nil, fmt.Errorf("parsing YAML file %s failed: %v", configPath, err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.Postgres == (postgres_client.Config{}) {
		return errors.New("postgres config not found")
	}

	if err := c.TokenGenerator.Validate(); err != nil {
		return fmt.Errorf("token generator config: %v", err)
	}

	if c.AlertSender == (alert_sender.Config{}) {
		return errors.New("alert sender config not found")
	}

	return nil
}
