package datadragon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	version = "15.12.1"
	baseURL = "https://ddragon.leagueoflegends.com"
)

var tempDir string

// Init creates a temporary directory and prefetches icons used by the frontend.
func Init() error {
	dir, err := os.MkdirTemp("", "datadragon-")
	if err != nil {
		return err
	}
	tempDir = dir

	if err := prefetchItems(); err != nil {
		return err
	}
	if err := prefetchSummoners(); err != nil {
		return err
	}
	if err := prefetchProfileIcons(); err != nil {
		return err
	}
	return prefetchRunes()
}

// Cleanup removes the downloaded icons.
func Cleanup() error {
	if tempDir == "" {
		return nil
	}
	return os.RemoveAll(tempDir)
}

// Get returns the local path of the requested icon. If the icon isn't cached yet,
// it will be downloaded once.
func Get(category, name string) (string, error) {
	if tempDir == "" {
		return "", errors.New("datadragon not initialised")
	}

	var local string
	switch category {
	case "item":
		local = filepath.Join(tempDir, "item", name)
	case "summoner":
		local = filepath.Join(tempDir, "summoner", name)
	case "profileicon":
		local = filepath.Join(tempDir, "profileicon", name)
	case "rune":
		local = filepath.Join(tempDir, "rune", filepath.FromSlash(name))
	default:
		return "", fmt.Errorf("unknown category %q", category)
	}

	if _, err := os.Stat(local); err == nil {
		return local, nil
	}
	if err := os.MkdirAll(filepath.Dir(local), 0o755); err != nil {
		return "", err
	}

	var url string
	switch category {
	case "item":
		url = fmt.Sprintf("%s/cdn/%s/img/item/%s", baseURL, version, name)
	case "summoner":
		url = fmt.Sprintf("%s/cdn/%s/img/spell/%s", baseURL, version, name)
	case "profileicon":
		url = fmt.Sprintf("%s/cdn/%s/img/profileicon/%s", baseURL, version, name)
	case "rune":
		url = fmt.Sprintf("%s/cdn/img/%s", baseURL, name)
	}
	if err := download(url, local); err != nil {
		return "", err
	}
	return local, nil
}

// -- internal helpers ---------------------------------------------------------

type itemData struct {
	Data map[string]struct {
		Image struct {
			Full string `json:"full"`
		} `json:"image"`
	} `json:"data"`
}

type summonerData struct {
	Data map[string]struct {
		Image struct {
			Full string `json:"full"`
		} `json:"image"`
	} `json:"data"`
}

type profileIconData struct {
	Data map[string]struct {
		Image struct {
			Full string `json:"full"`
		} `json:"image"`
	} `json:"data"`
}

type runeData []struct {
	Icon  string `json:"icon"`
	Slots []struct {
		Runes []struct {
			Icon string `json:"icon"`
		} `json:"runes"`
	} `json:"slots"`
}

func prefetchItems() error {
	url := fmt.Sprintf("%s/cdn/%s/data/en_US/item.json", baseURL, version)
	bts, err := fetch(url)
	if err != nil {
		return err
	}
	var data itemData
	if err := json.Unmarshal(bts, &data); err != nil {
		return err
	}

	for _, v := range data.Data {
		if _, err := Get("item", v.Image.Full); err != nil {
			return err
		}
	}
	return nil
}

func prefetchSummoners() error {
	url := fmt.Sprintf("%s/cdn/%s/data/en_US/summoner.json", baseURL, version)
	bts, err := fetch(url)
	if err != nil {
		return err
	}
	var data summonerData
	if err := json.Unmarshal(bts, &data); err != nil {
		return err
	}

	for _, v := range data.Data {
		if _, err := Get("summoner", v.Image.Full); err != nil {
			return err
		}
	}
	return nil
}

func prefetchProfileIcons() error {
	url := fmt.Sprintf("%s/cdn/%s/data/en_US/profileicon.json", baseURL, version)
	bts, err := fetch(url)
	if err != nil {
		return err
	}
	var data profileIconData
	if err := json.Unmarshal(bts, &data); err != nil {
		return err
	}

	for _, v := range data.Data {
		if _, err := Get("profileicon", v.Image.Full); err != nil {
			return err
		}
	}
	return nil
}

func prefetchRunes() error {
	url := fmt.Sprintf("%s/cdn/%s/data/en_US/runesReforged.json", baseURL, version)
	bts, err := fetch(url)
	if err != nil {
		return err
	}
	var data runeData
	if err := json.Unmarshal(bts, &data); err != nil {
		return err
	}

	for _, style := range data {
		if _, err := Get("rune", style.Icon); err != nil {
			return err
		}
		for _, slot := range style.Slots {
			for _, r := range slot.Runes {
				if _, err := Get("rune", r.Icon); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func download(url, dst string) error {
	bts, err := fetch(url)
	if err != nil {
		return err
	}
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bts)
	return err
}
