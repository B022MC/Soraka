package lcu

import (
	"Soraka/service/ddragon"
	"fmt"
	"log"
	"testing"
)

func TestGetCurrentSummoner(t *testing.T) {
	// 1. 获取 port + token
	port, token, err := GetLolClientApiInfo()
	if err != nil {
		t.Fatalf("GetLolClientApiInfo failed: %v", err)
	}

	// 2. 初始化 CLI 客户端（很关键）
	InitCli(port, token)

	// 3. 正常调用主函数
	info, err := WailsAPI{}.GetCurrentSummoner()
	if err != nil {
		t.Fatalf("GetCurrentSummoner returned error: %v", err)
	}

	// 4. 验证结果
	if info.DisplayName == "" {
		t.Error("DisplayName is empty")
	}
	if info.ProfileIconId == 0 {
		t.Error("ProfileIconId is 0")
	}
	if info.Level <= 0 {
		t.Errorf("Invalid Level: %d", info.Level)
	}
	if info.Avatar == "" {
		t.Error("Avatar is empty")
	}
	if info.TotalGames != info.Wins+info.Losses {
		t.Errorf("TotalGames mismatch: got %d, expected %d", info.TotalGames, info.Wins+info.Losses)
	}

	t.Logf("✅ SummonerInfo: %+v", info)
}
func TestListRecentMatchesSimple(t *testing.T) {
	port, token, err := GetLolClientApiInfo()
	if err != nil {
		t.Fatalf("GetLolClientApiInfo failed: %v", err)
	}

	InitCli(port, token)

	matches, err := ListRecentMatches(5)
	if err != nil {
		t.Fatalf("调用 ListRecentMatches 出错: %v", err)
	}

	for _, match := range matches {
		fmt.Printf("MatchID: %d | Result: %s | Mode: %s | Champion: %s\n",
			match.ID, match.Result, match.Mode, match.Champion)
		fmt.Printf("Spells: %v\n", match.Spells)
		fmt.Printf("Items: %v\n", match.Items)
		fmt.Printf("Map: %s\n", match.Map)
		fmt.Println("--------------------------")
	}
}
func TestGetLolClientApiInfo(t *testing.T) {
	// 获取 LCU 客户端 port + token
	port, token, err := GetLolClientApiInfo()
	if err != nil {
		log.Fatalf("获取 LCU 客户端信息失败: %v", err)
	}

	fmt.Printf("开始更新图标缓存 (port: %d)\n", port)

	// 调用 UpdateIconCache
	if err := ddragon.UpdateIconCache(port, token); err != nil {
		log.Fatalf("UpdateIconCache 失败: %v", err)
	}

	fmt.Println("✅ 图标缓存更新成功！")
}
