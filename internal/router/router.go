package router

import (
	"github.com/gin-gonic/gin"

	"Soraka/internal/handler"
)

// Register binds all routes to the gin engine.
func Register(r *gin.Engine) {
	r.POST("/config", handler.SaveConfig)
	r.GET("/config", handler.GetConfig)
}
