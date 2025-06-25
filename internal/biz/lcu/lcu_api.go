package lcu

import (
	"Soraka/internal/dal/lcu/models"
	lcuResp "Soraka/internal/dal/resp/lcu"
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

func (u *LcuApiUseCase) ListGamesByPUUID(puuid string, begin, limit int) (*models.GameListResp, error) {
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, begin, begin+limit))
	if err != nil {
		return nil, err
	}
	var data models.GameListResp
	if err := json.Unmarshal(bts, &data); err != nil {
		u.logger.Printf("获取比赛记录失败: %v", err)
		return nil, err
	}
	return &data, nil
}

func (u *LcuApiUseCase) QueryGameSummary(gameID int64) (*models.GameSummary, error) {
	_ = u.queryGameSummaryLimiter.Wait(context.Background())
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-match-history/v1/games/%d", gameID))
	if err != nil {
		return nil, err
	}
	var data models.GameSummary
	if err := json.Unmarshal(bts, &data); err != nil {
		return nil, err
	}
	if data.CommonResp.ErrorCode != "" {
		return nil, fmt.Errorf("查询对局详情失败: %s", data.CommonResp.Message)
	}
	return &data, nil
}

func (u *LcuApiUseCase) QuerySummonerByName(name string) (*models.Summoner, error) {
	bts, err := u.client.HttpGet(fmt.Sprintf("/lol-summoner/v1/summoners?name=%s", url.QueryEscape(name)))
	if err != nil {
		return nil, err
	}
	var data models.Summoner
	if err := json.Unmarshal(bts, &data); err != nil {
		u.logger.Printf("搜索用户失败: %v", err)
		return nil, err
	}
	if data.CommonResp.ErrorCode != "" {
		return nil, fmt.Errorf("搜索用户失败: %s", data.CommonResp.Message)
	}
	return &data, nil
}
func (u *LcuApiUseCase) GetCurrSummoner() (*models.CurrSummoner, error) {
	bts, err := u.client.HttpGet("/lol-summoner/v1/current-summoner")
	if err != nil {
		return nil, err
	}
	data := &models.CurrSummoner{}
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
				items = append(items, models.ItemIconURL(id))
			}
			spells := []string{
				models.SpellIconURL(int(p.Spell1Id)),
				models.SpellIconURL(int(p.Spell2Id)),
			}
			mb := &lcuResp.MatchBrief{
				ID:       g.GameId,
				Result:   result,
				Mode:     mapMode(g.GameMode),
				Kills:    p.Stats.Kills,
				Deaths:   p.Stats.Deaths,
				Assists:  p.Stats.Assists,
				CS:       p.Stats.TotalMinionsKilled + p.Stats.NeutralMinionsKilled,
				Gold:     p.Stats.GoldEarned,
				Time:     time.UnixMilli(g.GameCreation).Format("2006/01/02 15:04"),
				Level:    p.Stats.ChampLevel,
				Champion: models.ChampionIconURL(int(p.ChampionId)),
				Spells:   spells,
				Items:    items,
				Map:      mapMap(models.MapID(g.MapId)),
			}
			list = append(list, mb)
			break
		}
	}
	return list, nil
}

func mapMode(mode models.GameMode) string {
	switch mode {
	case models.GameModeARAM:
		return "极地大乱斗"
	case models.GameModeURF:
		return "无限乱斗"
	case models.GameModeClassic:
		return "匹配模式"
	default:
		return string(mode)
	}
}

func mapMap(id models.MapID) string {
	switch id {
	case models.MapIDClassic:
		return "召唤师峡谷"
	case models.MapIDARAM:
		return "极地大乱斗"
	case models.MapIDCheery:
		return "斗魂竞技场"
	default:
		return fmt.Sprintf("%d", id)
	}
}
