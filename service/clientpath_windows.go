//go:build windows

package service

import (
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// detectClientPath attempts to locate the client installation path via the Windows registry.
func detectClientPath() string {
	// Try Riot Client registry location first
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Riot Games\RiotClientInstalls`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		if path, _, err := k.GetStringValue("Riot Client"); err == nil {
			return path
		}
	}

	k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Riot Games\RiotClientInstalls`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		if path, _, err := k.GetStringValue("Riot Client"); err == nil {
			return path
		}
	}

	// If Riot Client was not found, try League of Legends entries
	paths := []string{
		`SOFTWARE\Riot Games\League of Legends`,
		`SOFTWARE\WOW6432Node\Riot Games\League of Legends`,
	}
	for _, p := range paths {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, p, registry.QUERY_VALUE)
		if err != nil {
			continue
		}
		// The key might contain different value names depending on the installer
		for _, name := range []string{"Path", "InstallLocation", "InstallDir"} {
			if val, _, err := k.GetStringValue(name); err == nil && val != "" {
				k.Close()
				return filepath.Join(val, "LeagueClient.exe")
			}
		}
		k.Close()
	}

	return ""
}
