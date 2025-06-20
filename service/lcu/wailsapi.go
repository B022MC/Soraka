package lcu

import (
	"os/exec"
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
	DisplayName   string `json:"displayName"`
	ProfileIconId int    `json:"profileIconId"`
	Region        string `json:"region"`
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
	info := SummonerInfo{
		DisplayName:   summoner.DisplayName,
		ProfileIconId: summoner.ProfileIconId,
		Region:        summoner.TagLine,
	}
	return info, nil
}

// StartClient attempts to start the League client if required.
func (WailsAPI) StartClient() error {
	path, err := FindLolPath()
	if err != nil {
		return err
	}
	cmd := exec.Command(path)
	return cmd.Start()
}
