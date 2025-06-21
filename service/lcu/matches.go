package lcu

import (
	"fmt"
	"strconv"

	"Soraka/dal/lcu/models"
)

const cdnBase = "https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default"

func championIcon(id int) string {
	if id <= 0 {
		return ""
	}
	return fmt.Sprintf("%s/v1/champion-icons/%d.png", cdnBase, id)
}

func spellIcon(id int) string {
	if id <= 0 {
		return ""
	}
	return fmt.Sprintf("%s/v1/summoner-spells/%d.png", cdnBase, id)
}

func itemIcon(id int) string {
	if id <= 0 {
		return ""
	}
	return fmt.Sprintf("%s/v1/items/%d.png", cdnBase, id)
}

var queueNameMap = map[models.GameQueueID]string{
	models.NormalQueueID:   "匹配模式",
	models.RankSoleQueueID: "单 / 双排",
	models.RankFlexQueueID: "灵活排位",
	models.ARAMQueueID:     "极地大乱斗",
	models.URFQueueID:      "无限火力",
	models.CheeryQueueID:   "斗魂竞技场",
}

var mapNameMap = map[models.MapID]string{
	models.MapIDClassic: "召唤师峡谷",
	models.MapIDARAM:    "极地大乱斗",
	models.MapIDCheery:  "斗魂竞技场",
}

type MatchHistory struct {
	ID       int64    `json:"id"`
	Result   string   `json:"result"`
	Mode     string   `json:"mode"`
	Kills    int      `json:"kills"`
	Deaths   int      `json:"deaths"`
	Assists  int      `json:"assists"`
	CS       int      `json:"cs"`
	Gold     int      `json:"gold"`
	Level    int      `json:"level"`
	Time     string   `json:"time"`
	Champion string   `json:"champion"`
	Spells   []string `json:"spells"`
	Items    []string `json:"items"`
	Map      string   `json:"map"`
}

func FetchRecentMatches(limit int) ([]MatchHistory, error) {
	if limit <= 0 || limit > 20 {
		limit = 10
	}

	summoner, err := GetCurrSummoner()
	if err != nil {
		return nil, err
	}

	resp, err := ListGamesByPUUID(summoner.Puuid, 0, limit)
	if err != nil {
		return nil, err
	}

	matches := make([]MatchHistory, 0, len(resp.Games.Games))
	for _, game := range resp.Games.Games {
		pid := 0
		for _, iden := range game.ParticipantIdentities {
			if iden.Player.SummonerId == summoner.SummonerId {
				pid = iden.ParticipantId
				break
			}
		}
		if pid == 0 {
			continue
		}
		found := false
		for _, p := range game.Participants {
			if p.ParticipantId != pid {
				continue
			}
			result := "lose"
			if p.Stats.Win {
				result = "win"
			}
			match := MatchHistory{
				ID:       game.GameId,
				Result:   result,
				Mode:     queueNameMap[models.GameQueueID(game.QueueId)],
				Kills:    p.Stats.Kills,
				Deaths:   p.Stats.Deaths,
				Assists:  p.Stats.Assists,
				CS:       p.Stats.TotalMinionsKilled,
				Gold:     p.Stats.GoldEarned,
				Level:    p.Stats.ChampLevel,
				Time:     game.GameCreationDate.Format("2006/01/02 15:04"),
				Champion: championIcon(p.ChampionId),
				Spells: []string{
					spellIcon(p.Spell1Id),
					spellIcon(p.Spell2Id),
				},
				Items: []string{
					itemIcon(p.Stats.Item0),
					itemIcon(p.Stats.Item1),
					itemIcon(p.Stats.Item2),
					itemIcon(p.Stats.Item3),
					itemIcon(p.Stats.Item4),
					itemIcon(p.Stats.Item5),
				},
				Map: mapNameMap[models.MapID(game.MapId)],
			}
			if match.Mode == "" {
				match.Mode = string(game.GameMode)
			}
			if match.Map == "" {
				match.Map = strconv.Itoa(game.MapId)
			}
			matches = append(matches, match)
			found = true
			break
		}
		if !found {
			continue
		}
	}
	return matches, nil
}
