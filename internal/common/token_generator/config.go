package token_generator

import "fmt"

type Config struct {
	SecretKey              string `yaml:"secret_key"`
	AccessTokenTTLMinutes  int    `yaml:"access_token_ttl_minutes"`
	RefreshTokenTTLMinutes int    `yaml:"refresh_token_ttl_minutes"`
}

func (c *Config) Validate() error {
	if c.SecretKey == "" {
		return fmt.Errorf("secret_key not provided")
	}

	if c.AccessTokenTTLMinutes == 0 {
		return fmt.Errorf("access_token_ttl_minutes not provided")
	}

	if c.RefreshTokenTTLMinutes == 0 {
		return fmt.Errorf("refresh_token_ttl_minutes not provided")
	}
	return nil
}
