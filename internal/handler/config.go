package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Soraka/internal/service"
)

type SaveConfigReq struct {
	Key string `json:"key" binding:"required"`
	Val string `json:"val" binding:"required"`
}

func SaveConfig(c *gin.Context) {
	var req SaveConfigReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.SaveConfig(req.Key, req.Val); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetConfig(c *gin.Context) {
	key := c.Query("key")
	val, err := service.GetConfig(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "val": val})
}
