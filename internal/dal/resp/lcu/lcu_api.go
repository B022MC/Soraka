package lcu

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
