package lcu

import (
	"fmt"
	"time"

	"Soraka/dal/lcu/models"
)

type MatchBrief struct {
	ID       int64    `json:"id"`
	Result   string   `json:"result"`
	Mode     string   `json:"mode"`
	Kills    int      `json:"kills"`
	Deaths   int      `json:"deaths"`
	Assists  int      `json:"assists"`
	CS       int      `json:"cs"`
	Gold     int      `json:"gold"`
	Time     string   `json:"time"`
	Level    int      `json:"level"`
	Champion string   `json:"champion"`
	Spells   []string `json:"spells"`
	Items    []string `json:"items"`
	Map      string   `json:"map"`
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

func ListRecentMatches(limit int) ([]MatchBrief, error) {
	summoner, err := GetCurrSummoner()
	if err != nil {
		return nil, err
	}
	if limit <= 0 || limit > 20 {
		limit = 10
	}
	resp, err := ListGamesByPUUID(summoner.Puuid, 0, limit)
	if err != nil {
		return nil, err
	}
	list := make([]MatchBrief, 0, len(resp.Games.Games))
	base := "http://localhost:8200/v1/lcu/proxy/lol-game-data/assets/v1"
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
				items = append(items, fmt.Sprintf("%s/items/%d.png", base, id))
			}
			spells := []string{
				fmt.Sprintf("%s/summoner-spells/%d.png", base, p.Spell1Id),
				fmt.Sprintf("%s/summoner-spells/%d.png", base, p.Spell2Id),
			}
			mb := MatchBrief{
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
				Champion: fmt.Sprintf("%s/champion-icons/%d.png", base, p.ChampionId),
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
