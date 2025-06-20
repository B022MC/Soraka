package lcu

import (
	"testing"
)

func TestGetLolClientApiInfoV3_Real(t *testing.T) {
	port, token, err := GetLolClientApiInfoV3()

	if err != nil {
		t.Logf("Expected: LeagueClientUx.exe not running or not matched, got error: %v", err)
		return // 如果没在运行，直接退出，不算失败
	}

	if port <= 0 {
		t.Errorf("Expected positive port, got %d", port)
	}
	if token == "" {
		t.Errorf("Expected non-empty token, got empty string")
	}

	t.Logf("Got port: %d, token: %s", port, token)
}
