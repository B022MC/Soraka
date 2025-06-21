import { defineStore } from "pinia";
import { setToken, getToken, clearToken } from "@/utils/auth";

export interface UserState {
  nickname?: string;
  avatar?: string;
  region?: string;
  tag?: string; // 召唤师标识 #12345
  rank?: string; // 当前段位
  winRate?: number; // 胜率（百分比）
  wins?: number; // 胜场
  losses?: number; // 败场
  totalGames?: number; // 总场次
  createtime: string;
  server: string;
  level?: number; // 等级
  xpSinceLastLevel?: number; // 当前等级已获得经验
  xpUntilNextLevel?: number; // 到下一级需要经验
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
