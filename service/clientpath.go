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
	cfg         clientConfig
	defaultCfg  clientConfig
	configFile  string
	defaultFile string
)

func init() {
	exe, err := os.Executable()
	if err != nil {
		exe = "."
	}
	baseDir := filepath.Dir(exe)
	configDir := filepath.Join(baseDir, "config")
	_ = os.MkdirAll(configDir, 0755)
	configFile = filepath.Join(configDir, "clientconfig.json")
	defaultFile = filepath.Join(configDir, "clientconfig.default.json")
	loadDefaultConfig()
	loadClientConfig()
}

func loadDefaultConfig() {
	data, err := os.ReadFile(defaultFile)
	if err == nil {
		_ = json.Unmarshal(data, &defaultCfg)
		cfg = defaultCfg
	}
}

func loadClientConfig() {
	data, err := os.ReadFile(configFile)
	if err == nil {
		_ = json.Unmarshal(data, &cfg)
	}
}

func saveClientConfig() {
	if cfg.Path == defaultCfg.Path || cfg.Path == "" {
		_ = os.Remove(configFile)
		return
	}
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
	path := getLoLPathByRegistry()
	if path != "" {
		cfg.Path = path
		saveClientConfig()
	}
	return path
}

// SaveClientPath allows the frontend to persist a user-specified path.
func SaveClientPath(path string) {
	cfg.Path = path
	saveClientConfig()
}
