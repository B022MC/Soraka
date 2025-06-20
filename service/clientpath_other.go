//go:build !windows

package service

// GetClientPath returns an empty string on non-Windows systems.
func detectClientPath() string {
	return ""
}
