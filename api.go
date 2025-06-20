package Soraka

import "github.com/gin-gonic/gin"

// Api defines HTTP handlers used by the frontend.
type Api struct{}

func (a *Api) DevHand(c *gin.Context) {
	c.JSON(200, gin.H{"message": "dev"})
}

// ProphetActiveMid is a placeholder middleware.
func (a *Api) ProphetActiveMid(c *gin.Context) {
	c.Next()
}

func (a *Api) QueryHorseBySummonerName(c *gin.Context) {
	c.JSON(200, gin.H{"message": "query horse by summoner"})
}

func (a *Api) GetAllConf(c *gin.Context) {
	c.JSON(200, gin.H{"config": "all"})
}

func (a *Api) UpdateClientConf(c *gin.Context) {
	c.JSON(200, gin.H{"message": "update ok"})
}

func (a *Api) GetLcuAuthInfo(c *gin.Context) {
	c.JSON(200, gin.H{"auth": "info"})
}

func (a *Api) GetAppInfo(c *gin.Context) {
	c.JSON(200, gin.H{"app": "info"})
}

func (a *Api) CopyHorseMsgToClipBoard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "copied"})
}

func (a *Api) LcuProxy(c *gin.Context) {
	c.JSON(200, gin.H{"message": "proxy"})
}
