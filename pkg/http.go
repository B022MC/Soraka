package pkg

import (
	"net/http"
	"slices"
	"strings"
)

// head field
const (
	HeadUserAgent   = "User-Agent"
	HeadContentType = "Content-Type"
)

// req path
const (
	FaviconReq       = "/favicon"
	VersionApi       = "/version"
	MetricsApi       = "/metrics"
	StatusApi        = "/status"
	DebugApiPrefix   = "/debug"
	SwaggerApiPrefix = "/swagger"
)

var (
	skipLogPathArr = []string{
		MetricsApi,
		VersionApi,
		StatusApi,
		FaviconReq,
	}
)

type (
	ServerInfoData struct {
		Timestamp   int64
		TimestampMs int64
	}
	VersionInfo struct {
		Version   string `json:"version" example:"1.0.0"`
		Commit    string `json:"commit" example:"github commit sha256"`
		BuildTime string `json:"buildTime" example:"2006-01-02 15:04:05"`
		BuildUser string `json:"buildUser" example:"buffge"`
	}
	IDReq struct {
		ID uint64 `json:"id" binding:"required,min=1"`
	}
	IDArrReq struct {
		IDArr []uint64 `json:"idArr" binding:"required,min=1,max=20"`
	}
	ListCommonReq struct {
		LastID string `json:"lastID" binding:""`
		Page   uint   `json:"page" binding:"omitempty,min=0"`
		Limit  uint   `json:"limit" binding:"required,min=1,max=20"`
	}
	ListCommonResp struct {
		LastID string `json:"lastID"`
		More   bool   `json:"more"`
		Count  int64  `json:"count"`
	}
)

func IsSkipLogReq(req *http.Request, statusCode int) bool {
	isDevApi := strings.Index(req.RequestURI, DebugApiPrefix) == 0 ||
		strings.Index(req.RequestURI, SwaggerApiPrefix) == 0
	return isDevApi || statusCode == http.StatusNotFound || req.Method == http.MethodOptions ||
		slices.Index(skipLogPathArr, req.RequestURI) >= 0
}

func GetUserAgent(r *http.Request) string {
	return GetUA(r)
}
func GetUA(r *http.Request) string {
	return r.Header.Get(HeadUserAgent)
}
func GetContentType(r *http.Request) string {
	return r.Header.Get(HeadContentType)
}
