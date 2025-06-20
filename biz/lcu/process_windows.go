//go:build windows

package lcu

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
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
func getPortTokenServerByPid(pid int) (string, string, string, error) {
	out, err := exec.Command("wmic", "process", "where", "processid="+strconv.Itoa(pid), "get", "commandline").Output()
	if err != nil {
		return "", "", "", err
	}
	s := string(out)
	rePort := regexp.MustCompile(`--app-port=([0-9]+)`)
	reToken := regexp.MustCompile(`--remoting-auth-token=([^\s"']+)`)
	reServer := regexp.MustCompile(`--rso_platform_id=([^\s"']+)`)
	port := ""
	token := ""
	server := ""
	if m := rePort.FindStringSubmatch(s); len(m) > 1 {
		port = m[1]
	}
	if m := reToken.FindStringSubmatch(s); len(m) > 1 {
		token = m[1]
	}
	if m := reServer.FindStringSubmatch(s); len(m) > 1 {
		server = m[1]
	}
	if port == "" || token == "" {
		return "", "", "", errors.New("unable to parse command line")
	}
	return port, token, server, nil
}
