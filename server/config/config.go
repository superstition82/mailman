package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

// Config is the configuration to start main server.
type Config struct {
	// Mode can be "prod" or "dev"
	Mode string `json:"mode"`
	// Port is the binding port for server
	Port int `json:"-"`
	// Data is the data directory
	Data string `json:"-"`
	// DSN points to where pocketmail stores its own data
	DSN string `json:"-"`
	// Version is the current version of server
	Version string `json:"version"`
}

func (p *Config) IsDev() bool {
	return p.Mode != "prod"
}

func checkDSN(dataDir string) (string, error) {
	// Convert to absolute path if relative path is supplied.
	if !filepath.IsAbs(dataDir) {
		relativeDir := filepath.Join(filepath.Dir(os.Args[0]), dataDir)
		absDir, err := filepath.Abs(relativeDir)
		if err != nil {
			return "", err
		}
		dataDir = absDir
	}

	// Trim trailing \ or / in case user supplies
	dataDir = strings.TrimRight(dataDir, "\\/")

	if _, err := os.Stat(dataDir); err != nil {
		return "", fmt.Errorf("unable to access data folder %s, err %w", dataDir, err)
	}

	return dataDir, nil
}

// Getconfig will return a config for dev or prod.
func Getconfig() (*Config, error) {
	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if config.Mode == "prod" && config.Data == "" {
		if runtime.GOOS == "windows" {
			config.Data = filepath.Join(os.Getenv("ProgramData"), "pocketmail")

			if _, err := os.Stat(config.Data); os.IsNotExist(err) {
				if err := os.MkdirAll(config.Data, 0770); err != nil {
					fmt.Printf("Failed to create data directory: %s, err: %+v\n", config.Data, err)
					return nil, err
				}
			}
		} else {
			config.Data = "/var/opt/pocketmail"
		}
	}

	dataDir, err := checkDSN(config.Data)
	if err != nil {
		fmt.Printf("Failed to check dsn: %s, err: %+v\n", dataDir, err)
		return nil, err
	}

	config.Data = dataDir
	dbFile := fmt.Sprintf("pocketmail_%s.db", config.Mode)
	config.DSN = filepath.Join(dataDir, dbFile)

	return &config, nil
}
