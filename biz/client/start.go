package client

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Start launches the League of Legends client if found.
func Start() error {
	base := GetClientPath()
	if base == "" {
		return fmt.Errorf("client path not configured")
	}
	for _, exe := range []string{"client.exe", "LeagueClient.exe"} {
		p := filepath.Join(base, exe)
		if _, err := os.Stat(p); err == nil {
			cmd := exec.Command(p)
			return cmd.Start()
		}
	}
	return fmt.Errorf("client executable not found")
}
