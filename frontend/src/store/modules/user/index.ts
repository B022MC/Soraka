import { defineStore } from "pinia";
import { setToken, getToken, clearToken } from "@/utils/auth";

export interface UserState {
  accountId?: number;               // 唯一账号ID
  summonerId?: number;              // 唯一召唤师ID
  puuid?: string;                   // 唯一玩家标识符
  nickname?: string;                // 游戏昵称（gameName / displayName）
  avatar?: string;                  // 头像完整URL
  region?: string;                  // 区域（如果后端有传）
  tag?: string;                     // 召唤师标识 #12345 (tagLine)
  rank?: string;                    // 当前段位（可选）
  winRate?: number;                 // 胜率（百分比，可选）
  wins?: number;                    // 胜场（可选）
  losses?: number;                  // 败场（可选）
  totalGames?: number;              // 总场次（可选）
  createtime?: string;              // 创建时间（可选）
  server?: string;                  // 服务器标识（可选，后端返回的话）
  level?: number;                   // 等级
  xpSinceLastLevel?: number;        // 当前等级已获得经验
  xpUntilNextLevel?: number;        // 到下一级需要经验
  percentCompleteForNextLevel?: number; // 升级进度百分比
  privacy?: string;                 // PUBLIC / PRIVATE
}


const useUserStore = defineStore("user", {
  state: (): UserState => ({
    nickname: undefined,
    avatar: undefined,
    region: "",
    tag: "",
    rank: "",
    winRate: undefined,
    server:"",
    wins: undefined,
    losses: undefined,
    totalGames: undefined,
    createtime: "",
    level: undefined,
    xpSinceLastLevel: undefined,
    xpUntilNextLevel: undefined,
  }),

  getters: {
    userInfo(state): UserState {
      return { ...state };
    },
  },

  actions: {
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial);
    },

    resetInfo() {
      this.$reset();
    },

    async setTokenArr(token: string) {
      setToken(token);
    },

    async info() {
      // const res = await getUserInfo()
      // this.setInfo(res)
    },

    async login() {
      setToken(undefined);
    },

    async logout(goLogin = false) {
      // TODO: 清空信息，跳转登录页等
    },

    clearloginfo() {
      clearToken();
    },
  },
});

export default useUserStore;
