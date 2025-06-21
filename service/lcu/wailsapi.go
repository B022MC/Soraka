package lcu

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

// WailsAPI exposes methods for the frontend via Wails service
// to interact with the local League client.
type WailsAPI struct{}

// Must implement Name() to satisfy application.Service interface
func (WailsAPI) Name() string { return "LcuApi" }

type AuthInfo struct {
	Port  int    `json:"port"`
	Token string `json:"token"`
}

type SummonerInfo struct {
	DisplayName      string  `json:"displayName"`
	ProfileIconId    int     `json:"profileIconId"`
	Region           string  `json:"region"`
	Avatar           string  `json:"avatar"`
	Tag              string  `json:"tag"`
	Rank             string  `json:"rank"`
	WinRate          float64 `json:"winRate"`
	Wins             int     `json:"wins"`
	Losses           int     `json:"losses"`
	TotalGames       int     `json:"totalGames"`
	Createtime       string  `json:"createtime"`
	Level            int     `json:"level"`
	XpSinceLastLevel int     `json:"xpSinceLastLevel"`
	XpUntilNextLevel int     `json:"xpUntilNextLevel"`
}

// GetClientPath returns the League of Legends executable path.
func (WailsAPI) GetClientPath() (string, error) {
	return FindLolPath()
}

// GetAuthInfo retrieves the LCU authentication information.
func (WailsAPI) GetAuthInfo() (AuthInfo, error) {
	port, token, err := GetLolClientApiInfo()
	if err != nil {
		return AuthInfo{}, err
	}
	return AuthInfo{Port: port, Token: token}, nil
}

// GetCurrentSummoner returns basic info about the current summoner.
func (WailsAPI) GetCurrentSummoner() (SummonerInfo, error) {
	summoner, err := GetCurrSummoner()
	if err != nil {
		return SummonerInfo{}, err
	}
	profile, err := GetSummonerProfile()
	if err != nil {
		return SummonerInfo{}, err
	}
	wins, _ := strconv.Atoi(profile.Lol.RankedWins)
	losses, _ := strconv.Atoi(profile.Lol.RankedLosses)
	total := wins + losses
	winRate := 0.0
	if total > 0 {
		winRate = float64(wins) * 100 / float64(total)
	}
	rank := profile.Lol.RankedLeagueTier
	if profile.Lol.RankedLeagueDivision != "" {
		rank = fmt.Sprintf("%s %s", rank, profile.Lol.RankedLeagueDivision)
	}
	port, _, _ := GetLolClientApiInfo()
	info := SummonerInfo{
		DisplayName:      summoner.DisplayName,
		ProfileIconId:    summoner.ProfileIconId,
		Region:           summoner.TagLine,
		Tag:              "#" + summoner.GameTag,
		Rank:             rank,
		WinRate:          winRate,
		Wins:             wins,
		Losses:           losses,
		TotalGames:       total,
		Createtime:       time.Now().Format(time.DateTime),
		Level:            summoner.SummonerLevel,
		XpSinceLastLevel: summoner.XpSinceLastLevel,
		XpUntilNextLevel: summoner.XpUntilNextLevel,
	}
	if port > 0 {
		info.Avatar = fmt.Sprintf("https://127.0.0.1:%d/lol-game-data/assets/v1/profile-icons/%d.jpg", port, summoner.ProfileIconId)
	}
	return info, nil
}

// StartClient attempts to start the League client if required.
func (WailsAPI) StartClient() error {
	base, err := FindLolPath()
	if err != nil {
		return err
	}

	// 常见的 LOL 客户端主程序名称
	executables := []string{
		"LeagueClient.exe",
		"client.exe",
	}

	for _, exe := range executables {
		fullPath := filepath.Join(base, exe)
		if _, err := os.Stat(fullPath); err == nil {
			cmd := exec.Command(fullPath)
			return cmd.Start()
		}
	}

	return fmt.Errorf("no executable client found in: %s", base)
}
