package lcu

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/pkg/errors"
)

type (
	RP struct {
		proxy *httputil.ReverseProxy
		Token string
		Port  int
	}
)

func NewRP(port int, token string) (*RP, error) {
	rawURL := fmt.Sprintf("https://127.0.0.1:%d", port)
	targetURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.Errorf("解析反向代理目标 URL 时出错: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host
		req.RequestURI = "" // 重要：必须清除 RequestURI，否则路径出错

		// 设置正确的 Basic Auth
		req.SetBasicAuth("riot", token)
	}

	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &RP{
		proxy: proxy,
		Token: token,
		Port:  port,
	}, nil
}
func (rp RP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.proxy.ServeHTTP(w, r)
}
