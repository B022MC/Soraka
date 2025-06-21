import { defHttp } from "@/utils/http";

export interface AuthInfo {
  port: number;
  token: string;
}

// 获取LCU登录信息
export function getLcuAuthInfo() {
  return defHttp.post<AuthInfo>({ url: "/v1/lcu/getAuthInfo" });
}

// 通过后端代理获取当前召唤师信息
export interface SummonerInfo {
  displayName: string;
  profileIconId: number;
  region: string;
  avatar: string;
  tag: string;
  rank: string;
  winRate: number;
  wins: number;
  losses: number;
  totalGames: number;
  createtime: string;
  level: number;
  xpSinceLastLevel: number;
  xpUntilNextLevel: number;
}

export function getCurrentSummoner() {
  return defHttp.get<SummonerInfo>({
    url: "/v1/lcu/proxy/lol-summoner/v1/current-summoner",
  });
}

// 启动LOL客户端
export function startClient() {
  return defHttp.post({ url: "/v1/client/start" });
}
