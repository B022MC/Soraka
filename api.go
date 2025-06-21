package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"

	ginApp "Soraka/pkg/gin"

	"Soraka/conf"
	"Soraka/dal/db/models"
	lcuModels "Soraka/dal/lcu/models"
	"Soraka/global"
	"Soraka/service/lcu"
)

type (
	Api struct {
		p *Prophet
	}
	summonerNameReq struct {
		SummonerName string `json:"summonerName"`
	}
	matchListReq struct {
		Limit int `json:"limit"`
	}
)

func (api Api) ProphetActiveMid(c *gin.Context) {
	app := ginApp.GetApp(c)
	if !api.p.LcuActive {
		app.ErrorMsg("请检查lol客户端是否已启动")
		return
	}
	c.Next()
}
func (api Api) QueryHorseBySummonerName(c *gin.Context) {
	app := ginApp.GetApp(c)
	d := &summonerNameReq{}
	if err := c.ShouldBind(d); err != nil {
		app.ValidError(err)
		return
	}
	summonerName := strings.TrimSpace(d.SummonerName)
	var summoner *lcuModels.Summoner
	if summonerName == "" {
		if api.p.CurrSummoner == nil {
			app.ErrorMsg("系统错误")
			return
		}
		summoner = lcu.ConvertCurrSummonerToSummoner(api.p.CurrSummoner)
	} else {
		info, err := lcu.QuerySummonerByName(summonerName)
		if err != nil || info.SummonerId <= 0 {
			app.ErrorMsg("未查询到召唤师")
			return
		}
		summoner = info
	}
	scoreInfo, err := GetUserScore(summoner)
	if err != nil {
		app.CommonError(err)
		return
	}
	scoreCfg := global.GetScoreConf()
	clientUserCfg := global.GetClientUserConf()
	var horse string
	for i, v := range scoreCfg.Horse {
		if scoreInfo.Score >= v.Score {
			horse = clientUserCfg.HorseNameConf[i]
			break
		}
	}
	app.Data(gin.H{
		"score":   scoreInfo.Score,
		"currKDA": scoreInfo.CurrKDA,
		"horse":   horse,
	})
}

func (api Api) CopyHorseMsgToClipBoard(c *gin.Context) {
	app := ginApp.GetApp(c)
	app.Success()
}
func (api Api) GetAllConf(c *gin.Context) {
	app := ginApp.GetApp(c)
	app.Data(global.GetClientUserConf())
}
func (api Api) UpdateClientConf(c *gin.Context) {
	app := ginApp.GetApp(c)
	d := &conf.UpdateClientUserConfReq{}
	if err := c.ShouldBind(d); err != nil {
		app.ValidError(err)
		return
	}
	cfg := global.SetClientUserConf(*d)
	bts, _ := json.Marshal(cfg)
	m := models.Config{}
	err := m.Update(models.LocalClientConfKey, string(bts))
	if err != nil {
		app.CommonError(err)
		return
	}
	app.Success()
}
func (api Api) DevHand(c *gin.Context) {
	app := ginApp.GetApp(c)
	app.Data(gin.H{
		"buffge": 23456,
	})
}
func (api Api) GetAppInfo(c *gin.Context) {
	app := ginApp.GetApp(c)
	app.Data(global.AppBuildInfo)
}
func (api Api) GetLcuAuthInfo(c *gin.Context) {
	app := ginApp.GetApp(c)
	port, token, err := lcu.GetLolClientApiInfo()
	if err != nil {
		app.CommonError(err)
		return
	}
	app.Data(gin.H{
		"token": token,
		"port":  port,
	})
}

func (api Api) GetClientPath(c *gin.Context) {
	app := ginApp.GetApp(c)
	path, err := lcu.FindLolPath()
	if err != nil {
		app.CommonError(err)
		return
	}
	app.Data(gin.H{"path": path})
}

func (api Api) StartClient(c *gin.Context) {
	app := ginApp.GetApp(c)
	path, err := lcu.FindLolPath()
	if err != nil {
		app.CommonError(err)
		return
	}
	if err = exec.Command(path).Start(); err != nil {
		app.CommonError(err)
		return
	}
	app.Success()
}

func (api Api) ListRecentMatches(c *gin.Context) {
	app := ginApp.GetApp(c)
	req := &matchListReq{}
	if err := c.ShouldBind(req); err != nil {
		app.ValidError(err)
		return
	}
	list, err := lcu.ListRecentMatches(req.Limit)
	if err != nil {
		app.CommonError(err)
		return
	}
	app.Data(list)
}
func (api Api) LcuProxy(c *gin.Context) {
	if api.p == nil {
		fmt.Println("[错误] api.p 是 nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LCU代理未初始化 [p]"})
		return
	}
	if api.p.lcuRP == nil {
		fmt.Println("[错误] api.p.lcuRP 是 nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LCU代理未初始化 [lcuRP]"})
		return
	}

	path := c.Param("any")
	c.Request.URL.Path = path
	api.p.lcuRP.ServeHTTP(c.Writer, c.Request)
}
