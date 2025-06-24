package lcu

import (
	"Soraka/pkg"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetCurrSummoner(t *testing.T) {
	// 假设你已有 logger
	logger := log.Default()

	// 假设你已获取到 token 和 port（可以手动指定或用 client info 查）
	port := 53000
	token := "TSqFYvVb3A3tgOaqULqWwQ"

	// 初始化 client 和 usecase
	client := pkg.NewClient(port, token)
	uc := NewLcuApiUseCase(logger, client)

	// 调用 GetCurrSummoner
	summoner, err := uc.GetCurrSummoner()
	if err != nil {
		t.Fatalf("GetCurrSummoner failed: %v", err)
	}

	// 打印结果（或断言检查值）
	bts, _ := json.MarshalIndent(summoner, "", "  ")
	fmt.Println(string(bts))
}
