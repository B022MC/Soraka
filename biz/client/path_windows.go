//go:build windows

package client

import (
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

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

	// 首字母大写
	if len(path) > 1 {
		path = strings.ToUpper(path[:1]) + path[1:]
	}
	return path
}
