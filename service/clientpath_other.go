//go:build !windows

package service

// detectClientPath returns an empty string on non-Windows systems.
func detectClientPath() string {
	return ""
}
