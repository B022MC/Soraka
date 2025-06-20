//go:build windows

package lcu

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"os/exec"
	"strconv"
	"strings"
)

// getTasklistPath checks available tasklist executable path.
func getTasklistPath() string {
	paths := []string{"tasklist", "C:/Windows/System32/tasklist.exe"}
	for _, p := range paths {
		cmd := exec.Command(p, "/FI", "imagename eq LeagueClientUx.exe", "/NH")
		if err := cmd.Run(); err == nil {
			return p
		}
	}
	return ""
}

// getLolClientPids returns all LeagueClientUx.exe PIDs.
func getLolClientPids(path string) ([]int, error) {
	out, err := exec.Command(path, "/FI", "imagename eq LeagueClientUx.exe", "/NH").Output()
	if err != nil {
		return nil, err
	}
	pids := []int{}
	if !bytes.Contains(out, []byte("LeagueClientUx.exe")) {
		return pids, nil
	}
	fields := bytes.Fields(out)
	for i, f := range fields {
		if bytes.Equal(f, []byte("LeagueClientUx.exe")) && i+1 < len(fields) {
			if pid, err := strconv.Atoi(string(fields[i+1])); err == nil {
				pids = append(pids, pid)
			}
		}
	}
	return pids, nil
}

// getLolClientPidsSlowly is a slower fallback using wmic.
func getLolClientPidsSlowly() ([]int, error) {
	out, err := exec.Command("wmic", "process", "where", "name='LeagueClientUx.exe'", "get", "ProcessId").Output()
	if err != nil {
		return nil, err
	}
	pids := []int{}
	fields := strings.Fields(string(out))
	for _, f := range fields {
		if pid, err := strconv.Atoi(f); err == nil {
			pids = append(pids, pid)
		}
	}
	return pids, nil
}

// isLolGameProcessExist checks if game process exists.
func isLolGameProcessExist(path string) bool {
	out, err := exec.Command(path, "/FI", "imagename eq League of Legends.exe", "/NH").Output()
	if err != nil {
		return false
	}
	return bytes.Contains(out, []byte("League of Legends.exe"))
}

// getPortTokenServerByPid retrieves LCU port, token and server from cmdline.

func getPortTokenServerByPid(pid int32) (string, string, string, error) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return "", "", "", err
	}
	cmds, err := proc.CmdlineSlice()
	if err != nil {
		return "", "", "", err
	}

	var port, token, server string
	for _, arg := range cmds {
		arg = strings.Trim(arg, `"`)
		if strings.HasPrefix(arg, "--app-port=") {
			port = strings.TrimPrefix(arg, "--app-port=")
		}
		if strings.HasPrefix(arg, "--remoting-auth-token=") {
			token = strings.TrimPrefix(arg, "--remoting-auth-token=")
		}
		if strings.HasPrefix(arg, "--rso_platform_id=") {
			server = strings.TrimPrefix(arg, "--rso_platform_id=")
		}
	}
	if port == "" || token == "" {
		return "", "", "", fmt.Errorf("missing port or token")
	}
	return port, token, server, nil
}
