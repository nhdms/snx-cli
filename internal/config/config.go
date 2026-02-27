package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server        string `yaml:"server"`
	Username      string `yaml:"username"`
	FixedPassword string `yaml:"fixed_password"`
	TOTPSecret    string `yaml:"totp_secret"`
}

func DefaultPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".snx-cli.yaml")
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Server == "" {
		return nil, fmt.Errorf("server is required in config")
	}
	if cfg.Username == "" {
		return nil, fmt.Errorf("username is required in config")
	}
	if cfg.FixedPassword == "" {
		return nil, fmt.Errorf("fixed_password is required in config")
	}
	if cfg.TOTPSecret == "" {
		return nil, fmt.Errorf("totp_secret is required in config")
	}

	return &cfg, nil
}

func CreateDefault(path string) error {
	cfg := Config{
		Server:        "vpn.example.com",
		Username:      "your-username",
		FixedPassword: "your-fixed-password",
		TOTPSecret:    "YOUR_BASE32_TOTP_SECRET",
	}

	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return err
	}

	return nil
}
