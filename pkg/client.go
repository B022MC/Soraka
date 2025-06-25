package pkg

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	httpCli = &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
	}

	cli *Client
)

type Client struct {
	port    int
	authPwd string
	baseURL string
}

func InitCli(port int, token string) {
	cli = NewClient(port, token)
}

func NewClient(port int, token string) *Client {
	return &Client{
		port:    port,
		authPwd: token,
		baseURL: GenerateClientApiURL(port, token),
	}
}

func (c *Client) HttpGet(path string) ([]byte, error) {
	return c.Req(http.MethodGet, path, nil)
}

func (c *Client) HttpPost(path string, body any) ([]byte, error) {
	return c.Req(http.MethodPost, path, body)
}

func (c *Client) HttpPatch(path string, body any) ([]byte, error) {
	return c.Req(http.MethodPatch, path, body)
}

func (c *Client) HttpDel(path string) ([]byte, error) {
	return c.Req(http.MethodDelete, path, nil)
}

func (c *Client) Req(method, path string, data any) ([]byte, error) {
	var body io.Reader
	if data != nil {
		bts, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("marshal request body failed: %w", err)
		}
		body = bytes.NewReader(bts)
	}

	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	// 可选：检查响应码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func GenerateClientApiURL(port int, token string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", httpScheme, authUserName, token, host, port)
}

const (
	authUserName = "riot"
	host         = "127.0.0.1"
	httpScheme   = "https"
)
