package global

import (
	"Soraka/internal/conf/app"
	"Soraka/internal/conf/client"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"

	"Soraka/internal/dal/lcu/models"
)

type (
	AppInfo struct {
		Version   string
		Commit    string
		BuildUser string
		BuildTime string
	}
	UserInfo struct {
		MacHash  string
		Summoner *models.SummonerProfileData
	}
)

// envKey
const (
	EnvKeyMode = "PROPHET_MODE"
)

const (
	ZapLoggerCleanupKey = "ZapLogger"
	LogWriterCleanupKey = "logWriter"
	OtelCleanupKey      = "otel"
)

var (
	cleanupsMu                      = &sync.Mutex{}
	defaultShouldAutoOpenBrowserCfg = true
	DefaultClientUserConf           = client.ClientUserConf{
		AutoAcceptGame:                 false,
		AutoPickChampID:                0,
		AutoBanChampID:                 0,
		AutoSendTeamHorse:              true,
		ShouldSendSelfHorse:            true,
		HorseNameConf:                  [6]string{"通天代", "小代", "上等马", "中等马", "下等马", "牛马"},
		ChooseSendHorseMsg:             [6]bool{true, true, true, true, true, true},
		ChooseChampSendMsgDelaySec:     3,
		ShouldInGameSaveMsgToClipBoard: true,
		ShouldAutoOpenBrowser:          &defaultShouldAutoOpenBrowserCfg,
	}
	DefaultAppConf = app.AppConf{
		CalcScore: app.CalcScoreConf{
			Enabled:            true,
			GameMinDuration:    900,
			AllowQueueIDList:   []int{430, 420, 450, 440, 1700},
			FirstBlood:         [2]float64{10, 5},
			PentaKills:         [1]float64{20},
			QuadraKills:        [1]float64{10},
			TripleKills:        [1]float64{5},
			JoinTeamRateRank:   [4]float64{10, 5, 5, 10},
			GoldEarnedRank:     [4]float64{10, 5, 5, 10},
			HurtRank:           [2]float64{10, 5},
			Money2hurtRateRank: [2]float64{10, 5},
			VisionScoreRank:    [2]float64{10, 5},
			MinionsKilled: [][2]float64{
				{10, 20},
				{9, 10},
				{8, 5},
			},
			KillRate: []app.RateItemConf{
				{Limit: 50, ScoreConf: [][2]float64{
					{15, 40},
					{10, 20},
					{5, 10},
				}},
				{Limit: 40, ScoreConf: [][2]float64{
					{15, 20},
					{10, 10},
					{5, 5},
				}},
			},
			HurtRate: []app.RateItemConf{
				{Limit: 40, ScoreConf: [][2]float64{
					{15, 40},
					{10, 20},
					{5, 10},
				}},
				{Limit: 30, ScoreConf: [][2]float64{
					{15, 20},
					{10, 10},
					{5, 5},
				}},
			},
			AssistRate: []app.RateItemConf{
				{Limit: 50, ScoreConf: [][2]float64{
					{20, 30},
					{18, 25},
					{15, 20},
					{10, 10},
					{5, 5},
				}},
				{Limit: 40, ScoreConf: [][2]float64{
					{20, 15},
					{15, 10},
					{10, 5},
					{5, 3},
				}},
			},
			AdjustKDA: [2]float64{2, 5},
			Horse: [6]app.HorseScoreConf{
				{Score: 180, Name: "通天代"},
				{Score: 150, Name: "小代"},
				{Score: 125, Name: "上等马"},
				{Score: 105, Name: "中等马"},
				{Score: 95, Name: "下等马"},
				{Score: 0.0001, Name: "牛马"},
			},
			MergeMsg: false,
		},
	}
	userInfo       = &UserInfo{}
	confMu         = sync.Mutex{}
	Conf           = new(app.AppConf)
	ClientUserConf = new(client.ClientUserConf)
	Logger         *zap.SugaredLogger
	Cleanups       = make(map[string]func(c context.Context) error)
	AppBuildInfo   = AppInfo{}
)

// DB
var (
	SqliteDB *gorm.DB
)

func GetUserInfo() UserInfo {
	confMu.Lock()
	defer confMu.Unlock()
	return *userInfo
}

func IsDevMode() bool {
	return GetEnv() == app.ModeDebug
}
func GetEnv() app.Mode {
	return Conf.Mode
}

func GetScoreConf() app.CalcScoreConf {
	confMu.Lock()
	defer confMu.Unlock()
	return Conf.CalcScore
}

func GetClientUserConf() client.ClientUserConf {
	confMu.Lock()
	defer confMu.Unlock()
	data := *ClientUserConf
	return data
}
func SetClientUserConf(cfg client.UpdateClientUserConfReq) *client.ClientUserConf {
	confMu.Lock()
	defer confMu.Unlock()
	if cfg.AutoAcceptGame != nil {
		ClientUserConf.AutoAcceptGame = *cfg.AutoAcceptGame
	}
	if cfg.AutoPickChampID != nil {
		ClientUserConf.AutoPickChampID = *cfg.AutoPickChampID
	}
	if cfg.AutoBanChampID != nil {
		ClientUserConf.AutoBanChampID = *cfg.AutoBanChampID
	}
	if cfg.AutoSendTeamHorse != nil {
		ClientUserConf.AutoSendTeamHorse = *cfg.AutoSendTeamHorse
	}
	if cfg.ShouldSendSelfHorse != nil {
		ClientUserConf.ShouldSendSelfHorse = *cfg.ShouldSendSelfHorse
	}
	if cfg.HorseNameConf != nil {
		ClientUserConf.HorseNameConf = *cfg.HorseNameConf
	}
	if cfg.ChooseSendHorseMsg != nil {
		ClientUserConf.ChooseSendHorseMsg = *cfg.ChooseSendHorseMsg
	}
	if cfg.ChooseChampSendMsgDelaySec != nil {
		ClientUserConf.ChooseChampSendMsgDelaySec = *cfg.ChooseChampSendMsgDelaySec
	}
	if cfg.ShouldInGameSaveMsgToClipBoard != nil {
		ClientUserConf.ShouldInGameSaveMsgToClipBoard = *cfg.ShouldInGameSaveMsgToClipBoard
	}
	if cfg.ShouldAutoOpenBrowser != nil {
		ClientUserConf.ShouldAutoOpenBrowser = cfg.ShouldAutoOpenBrowser
	}
	return ClientUserConf
}
func SetAppInfo(info AppInfo) {
	AppBuildInfo = info
}
