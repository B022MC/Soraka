import { defHttp } from "@/utils/http";

export interface AuthInfo {
  port: number;
  token: string;
}

export interface SummonerInfo {
  accountId: number;              // 唯一账号ID
  summonerId: number;             // 唯一召唤师ID
  puuid: string;                  // 唯一玩家标识符
  gameName: string;               // 游戏名
  tagLine: string;                // 游戏 tag
  displayName: string;            // 显示名（可能为空）
  profileIconId: number;          // 头像ID
  avatarUrl: string;              // 完整头像URL
  privacy: string;                // PUBLIC / PRIVATE
  summonerLevel: number;          // 等级
  xpSinceLastLevel: number;       // 当前等级已获得经验
  xpUntilNextLevel: number;       // 升级所需经验
  percentCompleteForNextLevel: number; // 升级进度百分比
  rerollPoints: {
    currentPoints: number;
    maxRolls: number;
    numberOfRolls: number;
    pointsCostToRoll: number;
    pointsToReroll: number;
  };
  unnamed: boolean;               // 是否未命名
  rank?: string;                  // 排名（需要另取）
  winRate?: number;               // 胜率（需要另取）
  wins?: number;                  // 胜场数（需要另取）
  losses?: number;                // 败场数（需要另取）
  totalGames?: number;            // 总对局数（需要另取）
  createtime?: string;            // 创建时间（需要另取）
}


export interface MatchItem {
  id: number;
  result: "win" | "lose";
  mode: string;
  kills: number;
  deaths: number;
  assists: number;
  cs: number;
  gold: number;
  time: string;
  level: number;
  champion: string;
  spells: string[];
  items: string[];
  map: string;
}

export function getRecentMatches(limit = 20) {
  return defHttp.post<MatchItem[]>({
    url: "/lcu/lcuApi/listRecentMatches",
    params: { limit },
  });
}
