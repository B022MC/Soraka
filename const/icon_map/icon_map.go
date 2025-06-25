package icon_map

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IconEntry struct {
	ID       int    `json:"id"`
	IconPath string `json:"iconPath"`
}

var (
	ItemIconMap  = make(map[int]string)
	ChampIconMap = make(map[int]string)
	SpellIconMap = make(map[int]string)
)

// NewIconMapDownloader 自动下载并初始化所有 icon map
func NewIconMapDownloader(itemURL, champURL, spellURL string) error {
	if err := downloadAndLoad(itemURL, ItemIconMap, "item"); err != nil {
		return err
	}
	if err := downloadAndLoad(champURL, ChampIconMap, "champion"); err != nil {
		return err
	}
	if err := downloadAndLoad(spellURL, SpellIconMap, "spell"); err != nil {
		return err
	}
	fmt.Println("所有图标 map 初始化完成")
	return nil
}

// 下载 JSON 并解析填充 map
func downloadAndLoad(url string, target map[int]string, logName string) error {
	fmt.Printf("[%s] 正在下载: %s\n", logName, url)
	if url == "" {
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("[%s] 下载失败: %w", logName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[%s] 下载返回非 200 状态: %d", logName, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[%s] 读取响应失败: %w", logName, err)
	}

	var entries []IconEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return fmt.Errorf("[%s] 解析 JSON 失败: %w", logName, err)
	}

	for _, e := range entries {
		target[e.ID] = e.IconPath
	}

	fmt.Printf("[%s] 加载 %d 个图标\n", logName, len(target))
	return nil
}

// 获取方法
func GetItemIcon(id int) (string, bool) {
	v, ok := ItemIconMap[id]
	return v, ok
}

func GetChampIcon(id int) (string, bool) {
	v, ok := ChampIconMap[id]
	return v, ok
}

func GetSpellIcon(id int) (string, bool) {
	v, ok := SpellIconMap[id]
	return v, ok
}
