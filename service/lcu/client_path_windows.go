//go:build windows

package lcu

import (
	"errors"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func FindLolPath() (string, error) {
	path := getLoLPathByRegistry()
	if path != "" {
		return path, nil
	}

	path, err := findLolPathFromUninstall()
	if err != nil {
		return "", err
	}
	return path, nil
}

// getLoLPathByRegistry 通过国服路径尝试查找 LOL 安装目录
func getLoLPathByRegistry() string {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Tencent\LOL`, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}
	defer key.Close()

	val, _, err := key.GetStringValue("InstallPath")
	if err != nil || val == "" {
		return ""
	}

	path := filepath.Join(val, "TCLS")
	path = filepath.Clean(path)
	path = strings.ReplaceAll(path, `\`, `/`)

	// 首字母大写（C:/ 变成 C:/）
	if len(path) > 1 {
		path = strings.ToUpper(path[:1]) + path[1:]
	}
	return path
}

// findLolPathFromUninstall 从常见的卸载表注册表中查找 LOL 安装位置
func findLolPathFromUninstall() (string, error) {
	roots := []registry.Key{registry.LOCAL_MACHINE, registry.CURRENT_USER}
	subkeys := []string{
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`,
		`SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`,
	}
	for _, root := range roots {
		for _, sub := range subkeys {
			k, err := registry.OpenKey(root, sub, registry.READ)
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
