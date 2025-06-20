package service

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// clientConfig stores persisted settings.
type clientConfig struct {
	Path string `json:"path"`
}

var (
	cfg        clientConfig
	configFile string
)

func init() {
	exe, err := os.Executable()
	if err == nil {
		configFile = filepath.Join(filepath.Dir(exe), "clientconfig.json")
	} else {
		configFile = "clientconfig.json"
	}
	loadClientConfig()
}

func loadClientConfig() {
	data, err := os.ReadFile(configFile)
	if err == nil {
		_ = json.Unmarshal(data, &cfg)
	}
}

func saveClientConfig() {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err == nil {
		_ = os.WriteFile(configFile, data, 0644)
	}
}

// GetClientPath returns the stored client path if present. If absent, it
// attempts to detect the path and persists the result.
func GetClientPath() string {
	if cfg.Path != "" {
		return cfg.Path
	}
	path := detectClientPath()
	if path != "" {
		cfg.Path = path
		saveClientConfig()
	}
	return path
}
