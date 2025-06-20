package lcu

import (
	"Soraka/biz/client"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getCredentials() (port, password string, err error) {
	base := client.GetClientPath()
	if base == "" {
		return "", "", fmt.Errorf("client path not configured")
	}
	data, err := os.ReadFile(filepath.Join(base, "lockfile"))
	if err != nil {
		return "", "", err
	}
	parts := strings.Split(string(data), ":")
	if len(parts) < 5 {
		return "", "", fmt.Errorf("invalid lockfile")
	}
	return parts[2], parts[3], nil
}

func call(port, pass, path string) bool {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%s%s", port, path), nil)
	if err != nil {
		return false
	}
	req.SetBasicAuth("riot", pass)
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr, Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	io.ReadAll(resp.Body)
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// CheckLogin concurrently checks if the LCU API is accessible.
func CheckLogin() bool {
	port, pass, err := getCredentials()
	if err != nil {
		return false
	}
	endpoints := []string{"/lol-summoner/v1/current-summoner", "/lol-login/v1/session"}
	concurrency := client.GetConcurrency()
	if concurrency <= 0 {
		concurrency = 1
	}

	jobs := make(chan string, len(endpoints))
	results := make(chan bool, len(endpoints))

	for w := 0; w < concurrency && w < len(endpoints); w++ {
		go func() {
			for ep := range jobs {
				results <- call(port, pass, ep)
			}
		}()
	}

	for _, ep := range endpoints {
		jobs <- ep
	}
	close(jobs)

	success := false
	for i := 0; i < len(endpoints); i++ {
		if <-results {
			success = true
		}
	}
	return success
}
