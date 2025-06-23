package proxy

import (
	"Soraka/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterLcuProxy 注册 LCU 代理路由
func RegisterLcuProxy(router gin.IRoutes, prophet *server.Prophet) {
	router.Any("/lcu/proxy/*any", func(c *gin.Context) {
		if prophet == nil || prophet.LcuRP == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "LCU代理未初始化"})
			return
		}
		path := c.Param("any")
		c.Request.URL.Path = path
		prophet.LcuRP.ServeHTTP(c.Writer, c.Request)
	})
}
