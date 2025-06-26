package lcu

import (
	lolConst "Soraka/const/lol"
	lcuResp "Soraka/internal/dal/resp/lcu"
	lcuModel "Soraka/internal/model/lcu"
	"Soraka/pkg"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

type LcuApiUseCase struct {
	logger                  *log.Logger
	queryGameSummaryLimiter *rate.Limiter
	client                  *pkg.Client
}

// NewLcuApiUseCase 构造函数
func NewLcuApiUseCase(logger *log.Logger, client *pkg.Client) *LcuApiUseCase {
	return &LcuApiUseCase{
		logger:                  logger,
		queryGameSummaryLimiter: rate.NewLimiter(rate.Every(time.Second/50), 50),
		client:                  client,
	}
}

func (u *LcuApiUseCase) ListGamesByPUUID(puuid string, begin, limit int) (*lcuModel.GameListResp, error) {
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begin, begin+limit))
	if err != nil {
		return nil, err
	}
	var data lcuModel.GameListResp
	if err := json.Unmarshal(bts, &data); err != nil {
		u.logger.Printf("获取比赛记录失败: %v", err)
		return nil, err
	}
	return &data, nil
}

func (u *LcuApiUseCase) QueryGameSummary(gameID int64) (*lcuModel.GameSummary, error) {
	_ = u.queryGameSummaryLimiter.Wait(context.Background())
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-match-history/v1/games/%d", gameID))
	if err != nil {
		return nil, err
	}
	var data lcuModel.GameSummary
	if err := json.Unmarshal(bts, &data); err != nil {
		return nil, err
	}
	if data.CommonResp.ErrorCode != "" {
		return nil, fmt.Errorf("查询对局详情失败: %s", data.CommonResp.Message)
	}
	return &data, nil
}

func (u *LcuApiUseCase) QuerySummonerByName(name string) (*lcuModel.Summoner, error) {
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-summoner/v1/summoners?name=%s", url.QueryEscape(name)))
	if err != nil {
		return nil, err
	}
	var data lcuModel.Summoner
	if err := json.Unmarshal(bts, &data); err != nil {
		u.logger.Printf("搜索用户失败: %v", err)
		return nil, err
	}
	if data.CommonResp.ErrorCode != "" {
		return nil, fmt.Errorf("搜索用户失败: %s", data.CommonResp.Message)
	}
	return &data, nil
}
func (u *LcuApiUseCase) GetCurrSummoner() (*lcuModel.CurrSummoner, error) {
	bts, err := u.client.HttpGet("/lol-summoner/v1/current-summoner")
	if err != nil {
		return nil, err
	}
	data := &lcuModel.CurrSummoner{}
	err = json.Unmarshal(bts, data)
	if err != nil {
		u.logger.Printf("获取当前召唤师失败 %v", zap.Error(err))
		return nil, err
	}
	if data.SummonerId == 0 {
		return nil, errors.New("获取当前召唤师失败")
	}

	if data.ProfileIconId > 0 {
		data.AvatarUrl = fmt.Sprintf("http://localhost:8200/lcu/proxy/lol-game-data/assets/v1/profile-icons/%d.jpg", data.ProfileIconId)
	}
	return data, nil
}
func (u *LcuApiUseCase) GetRankSummary() ([]*lcuResp.RankSummary, error) {
	bts, err := u.client.HttpGet("/lol-ranked/v1/current-ranked-stats")
	if err != nil {
		return nil, err
	}

	res := &lcuModel.RankedStatsResponse{}
	if err := json.Unmarshal(bts, res); err != nil {
		return nil, err
	}

	var result []*lcuResp.RankSummary
	for _, q := range res.Queues {
		var qType string
		switch q.QueueType {
		case "RANKED_SOLO_5x5":
			qType = "单 / 双排"
		case "RANKED_FLEX_SR":
			qType = "灵活排位"
		default:
			continue
		}

		total := q.Wins + q.Losses
		rate := "0%"
		if total > 0 {
			rate = fmt.Sprintf("%d%%", int(float64(q.Wins)/float64(total)*100))
		}

		item := &lcuResp.RankSummary{
			Type:  qType,
			Total: total,
			Rate:  rate,
			Wins:  q.Wins,
			Loses: q.Losses,
			Rank:  fmt.Sprintf("%s %s", mapTier(lolConst.RankTier(q.Tier)), q.Division),
			LP:    q.LeaguePoints,
			SeasonMax: fmt.Sprintf("%s %s",
				mapTier(lolConst.RankTier(q.HighestTier)),
				q.HighestDivision),
			LastSeason: fmt.Sprintf("%s %s",
				mapTier(lolConst.RankTier(q.PreviousSeasonEndTier)),
				q.PreviousSeasonEndDivision),
		}
		result = append(result, item)
	}

	return result, nil
}

func (u *LcuApiUseCase) AcceptGame() error {
	_, err := u.client.HttpPost("/lol-matchmaking/v1/ready-check/accept", nil)
	return err
}
func (u *LcuApiUseCase) ListRecentMatches(limit int) ([]*lcuResp.MatchBrief, error) {
	summoner, err := u.GetCurrSummoner()
	if err != nil {
		return nil, err
	}
	if limit <= 0 || limit > 20 {
		limit = 10
	}
	resp, err := u.ListGamesByPUUID(summoner.Puuid, 0, limit)
	if err != nil {
		return nil, err
	}
	list := make([]*lcuResp.MatchBrief, 0, len(resp.Games.Games))
	for _, g := range resp.Games.Games {
		partID := 0
		for _, p := range g.ParticipantIdentities {
			if p.Player.SummonerId == summoner.SummonerId {
				partID = p.ParticipantId
				break
			}
		}
		if partID == 0 {
			continue
		}
		for _, p := range g.Participants {
			if p.ParticipantId != partID {
				continue
			}
			result := "lose"
			if p.Stats.Win {
				result = "win"
			}
			items := []string{}
			it := []int{p.Stats.Item0, p.Stats.Item1, p.Stats.Item2, p.Stats.Item3, p.Stats.Item4, p.Stats.Item5}
			for _, id := range it {
				if id == 0 {
					continue
				}
				items = append(items, lolConst.ItemIconURL(id))
			}
			spells := []string{
				lolConst.SpellIconURL(int(p.Spell1Id)),
				lolConst.SpellIconURL(int(p.Spell2Id)),
			}
			mb := &lcuResp.MatchBrief{

				ID:         g.GameId,
				Result:     result,
				Mode:       mapMode(g.GameMode),
				ModeDetail: mapQueueId(g.QueueId),
				Kills:      p.Stats.Kills,
				Deaths:     p.Stats.Deaths,
				Assists:    p.Stats.Assists,
				CS:         p.Stats.TotalMinionsKilled + p.Stats.NeutralMinionsKilled,
				Gold:       p.Stats.GoldEarned,
				Time:       time.UnixMilli(g.GameCreation).Format("2006/01/02 15:04"),
				Level:      p.Stats.ChampLevel,
				Champion:   lolConst.ChampionIconURL(int(p.ChampionId)),
				Spells:     spells,
				Items:      items,
				Map:        mapMap(lolConst.MapID(g.MapId)),
			}

			list = append(list, mb)
			break
		}
	}
	return list, nil
}

func mapMode(mode lolConst.GameMode) string {
	switch mode {
	case lolConst.GameModeARAM:
		return "极地大乱斗"
	case lolConst.GameModeURF, lolConst.GameModeSnowURF:
		return "无限火力"
	case lolConst.GameModeClassic:
		return "匹配 / 排位"
	case lolConst.GameModeTFT:
		return "云顶之弈"
	case lolConst.GameModeCherry:
		return "斗魂竞技场"
	case lolConst.GameModeCustom:
		return "自定义模式"
	case lolConst.GameModeOneForAll:
		return "克隆大作战"
	case lolConst.GameModeAscension:
		return "奥德赛 / 飞升"
	case lolConst.GameModeKingPoro:
		return "帝企鹅之战"
	case lolConst.GameModeNexusBlitz:
		return "灵能激斗 / 极限闪击"
	case lolConst.GameModeUltimateSpellbook:
		return "奥术宝典"
	case lolConst.GameModeTutorial:
		return "教学模式"
	default:
		return string(mode)
	}
}

func mapMap(id lolConst.MapID) string {
	switch id {
	case lolConst.MapIDClassic:
		return "召唤师峡谷"
	case lolConst.MapIDARAM:
		return "极地大乱斗"
	case lolConst.MapIDCheery:
		return "斗魂竞技场"
	case lolConst.MapIDNexusBlitz:
		return "灵能激斗 / 极限闪击"
	case lolConst.MapIDTFT:
		return "云顶之弈"
	case lolConst.MapIDTutorial:
		return "教学地图"
	default:
		return fmt.Sprintf("%d", id)
	}
}
func mapQueueId(queueId lolConst.GameQueueID) string {
	switch queueId {
	case lolConst.NormalBlindQueueID:
		return "匹配（盲选）"
	case lolConst.NormalQueueID:
		return "匹配（自选）"
	case lolConst.RankSoloQueueID:
		return "单双排位"
	case lolConst.RankFlexQueueID:
		return "灵活排位"

	case lolConst.ARAMQueueID:
		return "极地大乱斗"
	case lolConst.URFQueueID, lolConst.SnowURFQueueID:
		return "无限火力"
	case lolConst.OneForAllQueueID:
		return "克隆大作战"
	case lolConst.AscensionQueueID:
		return "飞升模式"
	case lolConst.NexusBlitzQueueID:
		return "灵能激斗 / 极限闪击"
	case lolConst.CherryQueueID:
		return "斗魂竞技场"
	case lolConst.UltSpellbookQueueID:
		return "奥术宝典"

	case lolConst.BOTSimpleQueueID:
		return "人机入门"
	case lolConst.BOTNoviceQueueID:
		return "人机新手"
	case lolConst.BOTNormalQueueID:
		return "人机一般"

	case lolConst.TutorialQueueID:
		return "教学模式"
	case lolConst.PracticeToolQueueID:
		return "练习模式"

	default:
		return fmt.Sprintf("未知模式（%d）", queueId)
	}
}
func mapTier(tier lolConst.RankTier) string {
	switch tier {
	// 初阶段位
	case lolConst.RankTierIron:
		return "黑铁"
	case lolConst.RankTierBronze:
		return "青铜"
	case lolConst.RankTierSilver:
		return "白银"
	case lolConst.RankTierGold:
		return "黄金"
	case lolConst.RankTierPlatinum:
		return "白金"
	case lolConst.RankTierEmerald:
		return "翡翠"
	case lolConst.RankTierDiamond:
		return "钻石"

	// 高端段位
	case lolConst.RankTierMaster:
		return "大师"
	case lolConst.RankTierGrandMaster:
		return "宗师"
	case lolConst.RankTierChallenger:
		return "王者"

	default:
		return string(tier) // 保留原始值便于调试未知段位
	}
}
