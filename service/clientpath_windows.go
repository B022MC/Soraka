//go:build windows

package service

import (
	"golang.org/x/sys/windows/registry"
)

// detectClientPath attempts to locate the client installation path via the Windows registry.
func detectClientPath() string {
	// Try the 64bit location first
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Riot Games\RiotClientInstalls`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		if path, _, err := k.GetStringValue("Riot Client"); err == nil {
			return path
		}
	}
	// Fallback to Wow6432Node
	k, err = registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Riot Games\RiotClientInstalls`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()
		if path, _, err := k.GetStringValue("Riot Client"); err == nil {
			return path
		}
	}
	return ""
}
