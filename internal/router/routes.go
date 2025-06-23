package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, api ApiHandler) {
	r.Any("test", api.DevHand)
	initV1Module(r, api)
}

func initV1Module(r *gin.Engine, api ApiHandler) {
	v1 := r.Group("v1")
	v1.POST("horse/queryBySummonerName", api.ProphetActiveMid, api.QueryHorseBySummonerName)
	v1.POST("config/getAll", api.GetAllConf)
	v1.POST("config/update", api.UpdateClientConf)
	v1.POST("lcu/getAuthInfo", api.GetLcuAuthInfo)
	v1.POST("game/getPath", api.GetClientPath)
	v1.POST("game/start", api.StartClient)
	v1.POST("lcu/recentMatches", api.ListRecentMatches)
	v1.POST("app/getInfo", api.GetAppInfo)
	v1.POST("horse/copyHorseMsgToClipBoard", api.CopyHorseMsgToClipBoard)
	v1.Any("/lcu/proxy/*any", api.LcuProxy)
}
