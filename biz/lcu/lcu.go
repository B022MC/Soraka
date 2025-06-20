package lcu

import "fmt"

func getCredentials() (port, password string, err error) {
	path := getTasklistPath()
	var pids []int
	if path != "" {
		pids, err = getLolClientPids(path)
	} else {
		pids, err = getLolClientPidsSlowly()
	}
	if err != nil {
		return "", "", err
	}
	if len(pids) == 0 {
		return "", "", fmt.Errorf("client not running")
	}
	port, password, _, err = getPortTokenServerByPid(int32(pids[0]))
	if err != nil {
		return "", "", err
	}
	return port, password, nil
}

// GetCredentials exposes the current LCU port and auth token.
func GetCredentials() (string, string, error) {
	return getCredentials()
}

// CheckLogin returns login status and credentials.
// If both the port and auth token can be retrieved, the user is considered
// logged in. No network requests are made.
func CheckLogin() (bool, string, string) {
	port, pass, err := getCredentials()
	if err != nil || port == "" || pass == "" {
		return false, "", ""
	}
	return true, port, pass
}
