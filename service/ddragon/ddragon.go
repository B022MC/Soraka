package ddragon

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const iconBaseDir = "bin/icon"

const (
	authUserName = "riot"
	host         = "127.0.0.1"
	httpScheme   = "https"
)

func clientURL(port int, token string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", httpScheme, authUserName, token, host, port)
}

var httpCli = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: time.Second * 30,
}

func getGameVersion(port int, token string) (string, error) {
	url := fmt.Sprintf("%s/lol-patch/v1/game-version", clientURL(port, token))
	resp, err := httpCli.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var ver string
	if err := json.NewDecoder(resp.Body).Decode(&ver); err != nil {
		return "", err
	}
	if ver == "" {
		return "", fmt.Errorf("no version info")
	}
	return ver, nil
}

func UpdateIconCache(port int, token string) error {
	ver, err := getGameVersion(port, token)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(iconBaseDir, 0755); err != nil {
		return err
	}
	vf := filepath.Join(iconBaseDir, "version.txt")
	if b, err := os.ReadFile(vf); err == nil && string(b) == ver {
		return nil
	}
	if err := downloadItems(port, token); err != nil {
		return err
	}
	if err := downloadSummoners(port, token); err != nil {
		return err
	}
	if err := downloadProfileIcons(port, token); err != nil {
		return err
	}
	if err := downloadRunes(port, token); err != nil {
		return err
	}
	return os.WriteFile(vf, []byte(ver), 0644)
}
func downloadItems(port int, token string) error {
	url := fmt.Sprintf("%s/lol-game-data/assets/v1/items.json", clientURL(port, token))
	fmt.Printf("请求 items.json: %s\n", url)

	resp, err := httpCli.Get(url)
	if err != nil {
		return fmt.Errorf("请求 items.json 失败: %w", err)
	}
	defer resp.Body.Close()

	var data []struct {
		ID       int    `json:"id"`
		IconPath string `json:"iconPath"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("解析 items.json 失败: %w", err)
	}

	dir := filepath.Join(iconBaseDir, "item")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	for _, item := range data {
		if item.IconPath == "" {
			continue
		}
		src := fmt.Sprintf("%s%s", clientURL(port, token), item.IconPath)
		fileName := filepath.Base(item.IconPath)
		dst := filepath.Join(dir, fileName)

		fmt.Printf("下载: %s -> %s\n", src, dst)
		if err := downloadFile(src, dst); err != nil {
			fmt.Printf("⚠️ 下载失败: %v（继续）\n", err)
			continue
		}
	}
	fmt.Println("✅ item 图标下载完成")
	return nil
}

func downloadSummoners(port int, token string) error {
	url := fmt.Sprintf("%s/lol-game-data/assets/v1/summoner-spells.json", clientURL(port, token))
	resp, err := httpCli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var data struct {
		Data map[string]struct {
			Image struct {
				Full string `json:"full"`
			} `json:"image"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	dir := filepath.Join(iconBaseDir, "summoner")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	for _, v := range data.Data {
		if v.Image.Full == "" {
			continue
		}
		src := fmt.Sprintf("%s/lol-game-data/assets/v1/summoner-spells/%s", clientURL(port, token), v.Image.Full)
		dst := filepath.Join(dir, v.Image.Full)
		if err := downloadFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func downloadProfileIcons(port int, token string) error {
	url := fmt.Sprintf("%s/lol-game-data/assets/v1/profileicon.json", clientURL(port, token))
	resp, err := httpCli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var data struct {
		Data map[string]struct {
			Image struct {
				Full string `json:"full"`
			} `json:"image"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}
	dir := filepath.Join(iconBaseDir, "profileicon")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	for _, v := range data.Data {
		if v.Image.Full == "" {
			continue
		}
		src := fmt.Sprintf("%s/lol-game-data/assets/v1/profile-icons/%s", clientURL(port, token), v.Image.Full)
		dst := filepath.Join(dir, v.Image.Full)
		if err := downloadFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func downloadRunes(port int, token string) error {
	url := fmt.Sprintf("%s/lol-game-data/assets/v1/perkstyles.json", clientURL(port, token))
	resp, err := httpCli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var trees []struct {
		Icon  string `json:"icon"`
		Slots []struct {
			Runes []struct {
				Icon string `json:"icon"`
			} `json:"runes"`
		} `json:"slots"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&trees); err != nil {
		return err
	}
	dir := filepath.Join(iconBaseDir, "runes")
	for _, t := range trees {
		if t.Icon != "" {
			if err := saveRuneIcon(t.Icon, dir, port, token); err != nil {
				return err
			}
		}
		for _, s := range t.Slots {
			for _, r := range s.Runes {
				if r.Icon != "" {
					if err := saveRuneIcon(r.Icon, dir, port, token); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func saveRuneIcon(iconPath, base string, port int, token string) error {
	dst := filepath.Join(base, iconPath)
	src := fmt.Sprintf("%s/lol-game-data/assets/%s", clientURL(port, token), iconPath)
	return downloadFile(src, dst)
}

func downloadFile(url, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	resp, err := httpCli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func FindIconPath(name string) (string, error) {
	var result string
	err := filepath.Walk(iconBaseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == name {
			result = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", fmt.Errorf("icon %s not found", name)
	}
	return result, nil
}
