package lcu

import (
	lcuBiz "Soraka/internal/biz/lcu"
	lcuReq "Soraka/internal/dal/req/lcu"
	"Soraka/pkg"
	"Soraka/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type LcuApiService struct {
	uc *lcuBiz.LcuApiUseCase
}

// NewLcuApiService .
func NewLcuApiService() *LcuApiService {
	logger := log.Default()

	// 不强制退出，初始化时允许没有客户端
	port, token, err := lcuBiz.NewClientInfoUseCase().GetLolClientApiInfo()
	if err != nil {
		logger.Printf("[WARN] 初始化时未找到 LCU 客户端: %v", err)
		// 可以先返回 nil client，或留空，由调用方后续动态重建
		return &LcuApiService{
			uc: lcuBiz.NewLcuApiUseCase(logger, nil),
		}
	}

	client := pkg.NewClient(port, token)
	uc := lcuBiz.NewLcuApiUseCase(logger, client)

	return &LcuApiService{
		uc: uc,
	}
}

// 注册 Gin 路由
func (s *LcuApiService) RegisterGin(group gin.IRoutes) {
	group.GET("/lcuApi/curr_summoner", s.GetCurrSummoner)
	group.GET("/lcuApi/games_by_puuid", s.ListGamesByPUUID)
	group.GET("/lcuApi/summoner_by_name", s.QuerySummonerByName)
	group.GET("/lcuApi/listRecentMatches", s.ListRecentMatches)
	group.POST("/lcuApi/accept_game", s.AcceptGame)
}

// Wails 服务名称
func (s *LcuApiService) Name() string {
	return "LcuApiService"
}

func (s *LcuApiService) GetCurrSummoner(ctx *gin.Context) {
	data, err := s.uc.GetCurrSummoner()
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "获取成功", data)
}
func (s *LcuApiService) GetClientPath() (string, error) {
	return FindLolPath()
}
func (s *LcuApiService) ListGamesByPUUID(ctx *gin.Context) {
	puuid := ctx.Query("puuid")
	begin := ctx.DefaultQuery("begin", "0")
	limit := ctx.DefaultQuery("limit", "10")

	// 参数转换
	beginInt, _ := strconv.Atoi(begin)
	limitInt, _ := strconv.Atoi(limit)

	data, err := s.uc.ListGamesByPUUID(puuid, beginInt, limitInt)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "获取成功", data)
}

func (s *LcuApiService) QuerySummonerByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		response.ErrorMsg(ctx, "参数 name 不能为空")
		return
	}
	data, err := s.uc.QuerySummonerByName(name)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "获取成功", data)
}

func (s *LcuApiService) AcceptGame(ctx *gin.Context) {
	if err := s.uc.AcceptGame(); err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "已接受对局", "")
}
func (s *LcuApiService) ListRecentMatches(ctx *gin.Context) {
	req := &lcuReq.MatchListReq{}
	if err := ctx.ShouldBind(req); err != nil {
		return
	}
	list, err := s.uc.ListRecentMatches(req.Limit)
	if err != nil {
		response.ErrorMsg(ctx, err.Error())
		return
	}
	response.Ok(ctx, "获取成功", list)
}
