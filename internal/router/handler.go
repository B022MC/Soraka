// internal/router/handler.go
package router

import "github.com/gin-gonic/gin"

type ApiHandler interface {
	DevHand(*gin.Context)
	ProphetActiveMid(*gin.Context)
	QueryHorseBySummonerName(*gin.Context)
	GetAllConf(*gin.Context)
	UpdateClientConf(*gin.Context)
	GetLcuAuthInfo(*gin.Context)
	GetClientPath(*gin.Context)
	StartClient(*gin.Context)
	ListRecentMatches(*gin.Context)
	LcuProxy(*gin.Context)
	GetAppInfo(*gin.Context)
	CopyHorseMsgToClipBoard(*gin.Context)
}
