package ddragon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const iconBaseDir = "bin/icon"

func GetLatestVersion() (string, error) {
	resp, err := http.Get("https://ddragon.leagueoflegends.com/api/versions.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var vers []string
	if err := json.NewDecoder(resp.Body).Decode(&vers); err != nil {
		return "", err
	}
	if len(vers) == 0 {
		return "", fmt.Errorf("no version info")
	}
	return vers[0], nil
}

func UpdateIconCache() error {
	ver, err := GetLatestVersion()
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
	if err := downloadItems(ver); err != nil {
		return err
	}
	if err := downloadSummoners(ver); err != nil {
		return err
	}
	if err := downloadProfileIcons(ver); err != nil {
		return err
	}
	if err := downloadRunes(ver); err != nil {
		return err
	}
	return os.WriteFile(vf, []byte(ver), 0644)
}

func downloadItems(ver string) error {
	url := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/data/en_US/item.json", ver)
	resp, err := http.Get(url)
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
	dir := filepath.Join(iconBaseDir, "item")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	for _, v := range data.Data {
		if v.Image.Full == "" {
			continue
		}
		src := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/img/item/%s", ver, v.Image.Full)
		dst := filepath.Join(dir, v.Image.Full)
		if err := downloadFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func downloadSummoners(ver string) error {
	url := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/data/en_US/summoner.json", ver)
	resp, err := http.Get(url)
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
		src := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/img/spell/%s", ver, v.Image.Full)
		dst := filepath.Join(dir, v.Image.Full)
		if err := downloadFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func downloadProfileIcons(ver string) error {
	url := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/data/en_US/profileicon.json", ver)
	resp, err := http.Get(url)
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
		src := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%s", ver, v.Image.Full)
		dst := filepath.Join(dir, v.Image.Full)
		if err := downloadFile(src, dst); err != nil {
			return err
		}
	}
	return nil
}

func downloadRunes(ver string) error {
	url := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/data/en_US/runesReforged.json", ver)
	resp, err := http.Get(url)
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
			if err := saveRuneIcon(t.Icon, dir); err != nil {
				return err
			}
		}
		for _, s := range t.Slots {
			for _, r := range s.Runes {
				if r.Icon != "" {
					if err := saveRuneIcon(r.Icon, dir); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func saveRuneIcon(iconPath, base string) error {
	dst := filepath.Join(base, iconPath)
	src := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/img/%s", iconPath)
	return downloadFile(src, dst)
}

func downloadFile(url, dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	resp, err := http.Get(url)
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
