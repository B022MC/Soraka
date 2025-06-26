package lcu

import (
	lcuBiz "Soraka/internal/biz/lcu"
	lcuReq "Soraka/internal/dal/req/lcu"
	lcuResp "Soraka/internal/dal/resp/lcu"
	"Soraka/pkg"
	"Soraka/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type LcuApiService struct {
	uc       *lcuBiz.LcuApiUseCase
	clientUC *lcuBiz.ClientInfoUseCase
}

// NewLcuApiService 构建 LcuApiService 实例
func NewLcuApiService() *LcuApiService {
	logger := log.Default()

	clientUC := lcuBiz.NewClientInfoUseCase()

	port, token, err := clientUC.GetLolClientApiInfo()
	if err != nil {
		logger.Printf("[WARN] 初始化时未找到 LCU 客户端: %v", err)
		return &LcuApiService{
			uc:       lcuBiz.NewLcuApiUseCase(logger, nil),
			clientUC: clientUC,
		}
	}

	client := pkg.NewClient(port, token)
	uc := lcuBiz.NewLcuApiUseCase(logger, client)

	return &LcuApiService{
		uc:       uc,
		clientUC: clientUC,
	}
}

// 注册 Gin 路由
func (s *LcuApiService) RegisterGin(group gin.IRoutes) {
	group.GET("/lcuApi/curr_summoner", s.GetCurrSummoner)
	group.GET("/lcuApi/games_by_puuid", s.ListGamesByPUUID)
	group.GET("/lcuApi/summoner_by_name", s.QuerySummonerByName)
	group.POST("/lcuApi/accept_game", s.AcceptGame)
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
	return s.clientUC.FindLolPath()
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

//func (s *LcuApiService) ListRecentMatches(ctx *gin.Context) {
//	req := &lcuReq.MatchListReq{}
//	if err := ctx.ShouldBind(req); err != nil {
//		return
//	}
//	list, err := s.uc.ListRecentMatches(req.Limit)
//	if err != nil {
//		response.ErrorMsg(ctx, err.Error())
//		return
//	}
//	response.Ok(ctx, "获取成功", list)
//}

func (s *LcuApiService) StartClient() error {
	return s.clientUC.OpenClient()
}
func (s *LcuApiService) ListRecentMatches(req *lcuReq.MatchListReq) ([]*lcuResp.MatchBrief, error) {
	return s.uc.ListRecentMatches(req.Limit)
}
func (s *LcuApiService) GetRankSummary() ([]*lcuResp.RankSummary, error) {
	return s.uc.GetRankSummary()
}
