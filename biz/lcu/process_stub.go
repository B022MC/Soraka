//go:build !windows

package lcu

import "fmt"

func getTasklistPath() string { return "" }

func getLolClientPids(path string) ([]int, error) { return nil, fmt.Errorf("unsupported") }

func getLolClientPidsSlowly() ([]int, error) { return nil, fmt.Errorf("unsupported") }

func isLolGameProcessExist(path string) bool { return false }

func getPortTokenServerByPid(pid int) (string, string, string, error) {
	return "", "", "", fmt.Errorf("unsupported")
}
