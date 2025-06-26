package lol

import (
	"Soraka/const/icon_map"
	"fmt"
)

type (
	GameMode      string // 游戏模式
	GameQueueID   int    // 游戏队列模式id
	GameQueueType string // 游戏队列模式
	GameStatus    string // 游戏状态
	RankTier      string // 排位等级
	GameType      string // 游戏类型
	Spell         int    // 召唤师技能
	Champion      int    // 英雄
	Lane          string // 位置
	ChampionRole  string // 英雄角色
	GameFlow      string // 游戏状态
	MapID         int    // 地图id
	TeamID        int    // 队伍id
	TeamIDStr     string // 队伍id
	PlatformId    string // 大区id
)

const (
	GameWin  = "Win"
	GameFail = "Fail"
)

// 游戏模式
const (
	GameModeNone              GameMode = ""
	GameModeClassic           GameMode = "CLASSIC"      // 经典模式（匹配/排位）
	GameModeARAM              GameMode = "ARAM"         // 极地大乱斗
	GameModeTFT               GameMode = "TFT"          // 云顶之弈
	GameModeURF               GameMode = "URF"          // 无限火力
	GameModeCustom            GameMode = "PRACTICETOOL" // 自定义工具
	GameModeCherry            GameMode = "CHERRY"       // 斗魂竞技场
	GameModeOneForAll         GameMode = "ONEFORALL"    // 克隆大作战
	GameModeAscension         GameMode = "ASCENSION"    // 奥德赛/飞升
	GameModeKingPoro          GameMode = "KINGPORO"     // 帝企鹅之战
	GameModeNexusBlitz        GameMode = "NEXUSBLITZ"   // 灵能激斗/极限闪击
	GameModeTutorial          GameMode = "TUTORIAL"     // 教学模式
	GameModeSnowURF           GameMode = "SNOWURF"      // 雪地无限火力
	GameModeUltimateSpellbook GameMode = "ULTBOOK"      // 奥术宝典
)

// 队列模式
const (
	GameQueueTypeNormal   GameQueueType = "NORMAL"            // 匹配
	GameQueueTypeRankSolo GameQueueType = "RANKED_SOLO_5x5"   // 单双排
	GameQueueTypeRankFlex GameQueueType = "RANKED_FLEX_SR"    // 组排
	GameQueueTypeARAM     GameQueueType = "ARAM_UNRANKED_5x5" // 大乱斗5v5
	GameQueueTypeURF      GameQueueType = "URF"               // 无限火力
	GameQueueTypeBOT      GameQueueType = "BOT"               // 人机
	GameQueueTypeCustom   GameQueueType = "PRACTICETOOL"      // 自定义
)

// 游戏状态
const (
	GameStatusInQueue        GameStatus = "inQueue"                   // 队列中
	GameStatusInGame         GameStatus = "inGame"                    // 游戏中
	GameStatusChampionSelect GameStatus = "championSelect"            // 英雄选择中
	GameStatusOutOfGame      GameStatus = "outOfGame"                 // 退出游戏中
	GameStatusHostNormal     GameStatus = "hosting_NORMAL"            // 匹配组队中-队长
	GameStatusHostRankSolo   GameStatus = "hosting_RANKED_SOLO_5x5"   // 单排组队中-队长
	GameStatusHostRankFlex   GameStatus = "hosting_RANKED_FLEX_SR"    // 组排组队中-队长
	GameStatusHostARAM       GameStatus = "hosting_ARAM_UNRANKED_5x5" // 大乱斗5v5组队中-队长
	GameStatusHostURF        GameStatus = "hosting_URF"               // 无限火力组队中-队长
	GameStatusHostBOT        GameStatus = "hosting_BOT"               // 人机组队中-队长
)
const (
	GameFlowChampionSelect GameFlow = "ChampSelect" // 英雄选择中
	GameFlowReadyCheck     GameFlow = "ReadyCheck"  // 等待接受对局
	GameFlowInProgress     GameFlow = "InProgress"  // 进行中
	GameFlowMatchmaking    GameFlow = "Matchmaking" // 匹配中
	GameFlowNone           GameFlow = "None"        // 无
)

// 排位等级
const (
	RankTierIron        RankTier = "IRON"        // 黑铁
	RankTierBronze      RankTier = "BRONZE"      // 青铜
	RankTierSilver      RankTier = "SILVER"      // 白银
	RankTierGold        RankTier = "GOLD"        // 黄金
	RankTierPlatinum    RankTier = "PLATINUM"    // 白金
	RankTierEmerald     RankTier = "EMERALD"     // 翡翠
	RankTierDiamond     RankTier = "DIAMOND"     // 钻石
	RankTierMaster      RankTier = "MASTER"      // 大师
	RankTierGrandMaster RankTier = "GRANDMASTER" // 宗师
	RankTierChallenger  RankTier = "CHALLENGER"  // 王者
)

// 游戏类型
const (
	GameTypeMatch GameType = "MATCHED_GAME" // 匹配
)
const (
	// 匹配和排位
	NormalQueueID      GameQueueID = 430 // 匹配（自选模式）
	NormalBlindQueueID GameQueueID = 400 // 匹配（盲选）
	RankSoloQueueID    GameQueueID = 420 // 单双排位
	RankFlexQueueID    GameQueueID = 440 // 灵活组排

	// 特殊模式
	ARAMQueueID         GameQueueID = 450  // 极地大乱斗
	URFQueueID          GameQueueID = 900  // 无限火力（随机URF）
	SnowURFQueueID      GameQueueID = 920  // 雪地无限火力
	OneForAllQueueID    GameQueueID = 1020 // 克隆大作战
	AscensionQueueID    GameQueueID = 910  // 飞升模式
	NexusBlitzQueueID   GameQueueID = 1300 // 灵能激斗 / 极限闪击
	CherryQueueID       GameQueueID = 1700 // 斗魂竞技场
	UltSpellbookQueueID GameQueueID = 1400 // 奥术宝典

	// 人机
	BOTSimpleQueueID GameQueueID = 830 // 人机入门
	BOTNoviceQueueID GameQueueID = 840 // 人机新手
	BOTNormalQueueID GameQueueID = 850 // 人机一般

	// 教学 / 自定义
	TutorialQueueID     GameQueueID = 2000 // 教学模式
	PracticeToolQueueID GameQueueID = 2010 // 练习模式 / 自定义工具
)

// 地图id
const (
	MapIDClassic    MapID = 11 // 召唤师峡谷
	MapIDARAM       MapID = 12 // 极地大乱斗
	MapIDCheery     MapID = 30 // 斗魂竞技场
	MapIDNexusBlitz MapID = 21 // 灵能激斗 / 极限闪击
	MapIDTFT        MapID = 25 // 云顶之弈
	MapIDTutorial   MapID = 3  // 教学地图
)

// 队伍id
const (
	TeamIDNone    TeamID    = 0     // 未知
	TeamIDBlue    TeamID    = 100   // 蓝色方
	TeamIDRed     TeamID    = 200   // 红色方
	TeamIDStrNone TeamIDStr = ""    // 未知
	TeamIDStrBlue TeamIDStr = "100" // 蓝色方
	TeamIDStrRed  TeamIDStr = "200" // 红色方
)

// 平台大区 ID（用于 profile.PlatformId）
const (
	PlatformIDHN1  = "HN1"  // 艾欧尼亚
	PlatformIDHN2  = "HN2"  // 祖安
	PlatformIDHN3  = "HN3"  // 诺克萨斯
	PlatformIDHN4  = "HN4"  // 班德尔城
	PlatformIDHN5  = "HN5"  // 皮尔特沃夫
	PlatformIDHN6  = "HN6"  // 战争学院
	PlatformIDHN7  = "HN7"  // 巨神峰
	PlatformIDHN8  = "HN8"  // 雷瑟守备
	PlatformIDHN9  = "HN9"  // 裁决之地
	PlatformIDHN10 = "HN10" // 黑色玫瑰
	PlatformIDHN11 = "HN11" // 暗影岛
	PlatformIDHN12 = "HN12" // 钢铁烈阳
	PlatformIDHN13 = "HN13" // 水晶之痕
	PlatformIDHN14 = "HN14" // 均衡教派
	PlatformIDHN15 = "HN15" // 影流
	PlatformIDHN16 = "HN16" // 守望之海
	PlatformIDHN17 = "HN17" // 征服之海
	PlatformIDHN18 = "HN18" // 卡拉曼达
	PlatformIDHN19 = "HN19" // 皮城警备
	PlatformIDHN20 = "HN20" // 巨龙之巢
)

var PlatformMap = map[string]string{
	PlatformIDHN1:  "艾欧尼亚",
	PlatformIDHN2:  "祖安",
	PlatformIDHN3:  "诺克萨斯",
	PlatformIDHN4:  "班德尔城",
	PlatformIDHN5:  "皮尔特沃夫",
	PlatformIDHN6:  "战争学院",
	PlatformIDHN7:  "巨神峰",
	PlatformIDHN8:  "雷瑟守备",
	PlatformIDHN9:  "裁决之地",
	PlatformIDHN10: "黑色玫瑰",
	PlatformIDHN11: "暗影岛",
	PlatformIDHN12: "钢铁烈阳",
	PlatformIDHN13: "水晶之痕",
	PlatformIDHN14: "均衡教派",
	PlatformIDHN15: "影流",
	PlatformIDHN16: "守望之海",
	PlatformIDHN17: "征服之海",
	PlatformIDHN18: "卡拉曼达",
	PlatformIDHN19: "皮城警备",
	PlatformIDHN20: "巨龙之巢",
}

// 召唤师技能
const (
	SpellPingZhang Spell = 21 // 屏障
	SpellShanXian  Spell = 4  // 闪现
)

// 位置
const (
	LaneTop    Lane = "JUNGLE" // 上路
	LaneJungle Lane = "JUNGLE" // 打野
	LaneMiddle Lane = "MIDDLE" // 中路
	LaneBottom Lane = "BOTTOM" // 下路
)

// 英雄角色
const (
	ChampionRoleSolo    ChampionRole = "SOLE"        // 单人路
	ChampionRoleSupport ChampionRole = "DUO_SUPPORT" // 辅助
	ChampionRoleADC     ChampionRole = "DUO_CARRY"   // adc
	ChampionRoleNone    ChampionRole = "NONE"        // 无 一般是打野
)

// 游戏大区
const (
	PlatformIdHN1 PlatformId = "HN1"
)

// LCU 静态资源路径模板
const (
	LCUAssetBase = "http://localhost:8200/lcu/proxy%s"
	ItemIconTpl  = LCUAssetBase + "/items/%s"                                                      // 正确路径
	SpellIconTpl = LCUAssetBase + "/summoner-spells/%s"                                            // 正确路径
	ChampIconTpl = "http://localhost:8200/lcu/proxy/lol-game-data/assets/v1/champion-icons/%d.png" // 正确路径
	MapIconTpl   = LCUAssetBase + "/map-icons/%d"                                                  // 如果你有地图图标
)

// 构建物品图标 URL
func ItemIconURL(itemID int) string {
	if itemID <= 0 {
		return ""
	}
	filename, ok := icon_map.ItemIconMap[itemID]
	if !ok {
		return ""
	}
	return fmt.Sprintf(LCUAssetBase, filename)
}

// 构建召唤师技能图标 URL
func SpellIconURL(spellID int) string {
	if spellID <= 0 {
		return ""
	}
	return fmt.Sprintf(LCUAssetBase, icon_map.SpellIconMap[spellID])
}

// 构建英雄图标 URL
func ChampionIconURL(champID int) string {
	if champID <= 0 {
		return ""
	}
	return fmt.Sprintf(ChampIconTpl, champID)
}

// 构建地图图标 URL（如需要）
func MapIconURL(mapID int) string {
	if mapID <= 0 {
		return ""
	}
	return fmt.Sprintf(LCUAssetBase, mapID)
}
