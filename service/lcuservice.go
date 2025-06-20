package service

import (
	"Soraka/biz/lcu"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LcuService exposes LCU related helpers.
type LcuService struct{}

// CheckLogin returns the login status along with LCU credentials.
func (LcuService) CheckLogin() (bool, string, string) {
	return lcu.CheckLogin()
}

func (LcuService) GetCredentials() (string, string, error) {
	return lcu.GetCredentials()
}

// UserInfo represents basic player information returned from the LCU.
type UserInfo struct {
	DisplayName   string `json:"displayName"`
	ProfileIconID int    `json:"profileIconId"`
	Region        string `json:"region"`
	Port          string `json:"port"`
}

// GetCurrentUserInfo fetches current summoner info and region from the LCU.
func (LcuService) GetCurrentUserInfo() (UserInfo, error) {
	var info UserInfo
	port, token, err := lcu.GetCredentials()
	if err != nil {
		return info, err
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%s/lol-summoner/v1/current-summoner", port), nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("riot:"+token)))
	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return info, err
	}

	reqReg, _ := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%s/riotclient/region-locale", port), nil)
	reqReg.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("riot:"+token)))
	respReg, err := client.Do(reqReg)
	if err == nil {
		defer respReg.Body.Close()
		var reg struct {
			Region string `json:"region"`
		}
		if err := json.NewDecoder(respReg.Body).Decode(&reg); err == nil {
			info.Region = reg.Region
		}
	}

	info.Port = port
	return info, nil
}
