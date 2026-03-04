package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules             RulesConfig `yaml:"rules"`
	SensitivePatterns []string    `yaml:"sensitive_patterns"`
}

type RulesConfig struct {
	Lowercase       bool `yaml:"lowercase"`
	English         bool `yaml:"english"`
	NoSpecialChars  bool `yaml:"no_special_chars"`
	NoSensitiveData bool `yaml:"no_sensitive_data"`
}

func DefaultConfig() *Config {
	return &Config{
		Rules: RulesConfig{
			Lowercase:       true,
			English:         true,
			NoSpecialChars:  true,
			NoSensitiveData: true,
		},
	}
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
