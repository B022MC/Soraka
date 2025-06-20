//go:build windows

package lcu

import (
	"errors"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// FindLolPath attempts to locate the League of Legends executable path from the Windows registry.
func FindLolPath() (string, error) {
	roots := []registry.Key{registry.LOCAL_MACHINE, registry.CURRENT_USER}
	subkeys := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}
	for _, r := range roots {
		for _, sub := range subkeys {
			k, err := registry.OpenKey(r, sub, registry.READ)
			if err != nil {
				continue
			}
			names, err := k.ReadSubKeyNames(-1)
			if err != nil {
				k.Close()
				continue
			}
			for _, name := range names {
				sk, err := registry.OpenKey(k, name, registry.READ)
				if err != nil {
					continue
				}
				display, _, _ := sk.GetStringValue("DisplayName")
				if strings.Contains(display, "英雄联盟") || strings.Contains(display, "League of Legends") {
					if loc, _, err := sk.GetStringValue("InstallLocation"); err == nil && loc != "" {
						sk.Close()
						k.Close()
						return filepath.Join(loc, "League of Legends.exe"), nil
					}
					if icon, _, err := sk.GetStringValue("DisplayIcon"); err == nil && icon != "" {
						sk.Close()
						k.Close()
						return icon, nil
					}
				}
				sk.Close()
			}
			k.Close()
		}
	}
	return "", errors.New("LOL path not found in registry")
}
