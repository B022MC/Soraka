//go:build windows

package lcu

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"Soraka/pkg/windows/process"
	"golang.org/x/sys/windows/registry"
)

const (
	lolUxProcessName = "LeagueClientUx.exe"
)

var (
	lolCommandlineReg     = regexp.MustCompile(`--remoting-auth-token=(.+?)" ".*?--app-port=(\d+)"`)
	ErrLolProcessNotFound = errors.New("没有找到lol客户端进程")
)

type ClientInfoUseCase struct{}

func NewClientInfoUseCase() *ClientInfoUseCase {
	return &ClientInfoUseCase{}
}

// 获取 LCU API 信息
func (uc *ClientInfoUseCase) GetLolClientApiInfo() (port int, token string, err error) {
	cmdline, err := process.GetProcessCommand(lolUxProcessName)
	if err != nil {
		return 0, "", ErrLolProcessNotFound
	}

	btsChunk := lolCommandlineReg.FindSubmatch([]byte(cmdline))
	if len(btsChunk) < 3 {
		return 0, "", ErrLolProcessNotFound
	}

	token = string(btsChunk[1])
	port, err = strconv.Atoi(string(btsChunk[2]))
	return
}

// FindLolPath 查找 LOL 客户端可执行文件路径
func (uc *ClientInfoUseCase) FindLolPath() (string, error) {
	basePath := uc.getLoLPathByRegistry()
	if basePath != "" {
		if exePath := uc.appendClientExe(basePath); exePath != "" {
			return exePath, nil
		}
	}

	basePath, err := uc.findLolPathFromUninstall()
	if err != nil {
		return "", err
	}
	if exePath := uc.appendClientExe(basePath); exePath != "" {
		return exePath, nil
	}

	return "", errors.New("未找到 LOL 客户端可执行文件")
}

// getLoLPathByRegistry 通过国服路径尝试查找 LOL 安装目录
func (uc *ClientInfoUseCase) getLoLPathByRegistry() string {
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

	if len(path) > 1 {
		path = strings.ToUpper(path[:1]) + path[1:]
	}
	return path
}

// findLolPathFromUninstall 从常见的卸载表注册表中查找 LOL 安装位置
func (uc *ClientInfoUseCase) findLolPathFromUninstall() (string, error) {
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

// OpenClient 打开 LOL 客户端
func (uc *ClientInfoUseCase) OpenClient() error {
	path, err := uc.FindLolPath()
	if err != nil {
		return err
	}

	cmd := exec.Command(path)
	return cmd.Start()
}

// appendClientExe 尝试在路径上拼接常见可执行文件名，并检查存在性
func (uc *ClientInfoUseCase) appendClientExe(basePath string) string {
	candidates := []string{
		filepath.Join(basePath, "LeagueClient.exe"),
		filepath.Join(basePath, "Client.exe"),
	}

	for _, path := range candidates {
		if fileExists(path) {
			return path
		}
	}
	return ""
}

// fileExists 检查文件是否存在且不是目录
func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
